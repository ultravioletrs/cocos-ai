// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package grpc

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/absmach/magistrala/pkg/errors"
	"github.com/ultravioletrs/cocos/agent"
	"github.com/ultravioletrs/cocos/agent/cvm"
	"github.com/ultravioletrs/cocos/agent/cvm/server"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

var (
	errCorruptedManifest = errors.New("received manifest may be corrupted")
	sendTimeout          = 5 * time.Second
)

type CVMClient struct {
	mu            sync.Mutex
	stream        cvm.CVMService_ProcessClient
	svc           agent.Service
	messageQueue  chan *cvm.ClientStreamMessage
	logger        *slog.Logger
	runReqManager *runRequestManager
	sp            *server.AgentServer
}

// NewClient returns new gRPC client instance.
func NewClient(stream cvm.CVMService_ProcessClient, svc agent.Service, messageQueue chan *cvm.ClientStreamMessage, logger *slog.Logger, sp *server.AgentServer) CVMClient {
	return CVMClient{
		stream:        stream,
		svc:           svc,
		messageQueue:  messageQueue,
		logger:        logger,
		runReqManager: newRunRequestManager(),
		sp:            sp,
	}
}

func (client *CVMClient) Process(ctx context.Context, cancel context.CancelFunc) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return client.handleIncomingMessages(ctx)
	})

	eg.Go(func() error {
		return client.handleOutgoingMessages(ctx)
	})

	return eg.Wait()
}

func (client *CVMClient) handleIncomingMessages(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			req, err := client.stream.Recv()
			if err != nil {
				return err
			}
			if err := client.processIncomingMessage(ctx, req); err != nil {
				return err
			}
		}
	}
}

func (client *CVMClient) processIncomingMessage(ctx context.Context, req *cvm.ServerStreamMessage) error {
	switch mes := req.Message.(type) {
	case *cvm.ServerStreamMessage_RunReqChunks:
		return client.handleRunReqChunks(ctx, mes)
	case *cvm.ServerStreamMessage_StopComputation:
		go client.handleStopComputation(ctx, mes)
	default:
		return errors.New("unknown message type")
	}
	return nil
}

func (client *CVMClient) handleRunReqChunks(ctx context.Context, mes *cvm.ServerStreamMessage_RunReqChunks) error {
	buffer, complete := client.runReqManager.addChunk(mes.RunReqChunks.Id, mes.RunReqChunks.Data, mes.RunReqChunks.IsLast)

	if complete {
		var runReq cvm.ComputationRunReq
		if err := proto.Unmarshal(buffer, &runReq); err != nil {
			return errors.Wrap(err, errCorruptedManifest)
		}

		go client.executeRun(ctx, &runReq)
	}

	return nil
}

func (client *CVMClient) executeRun(ctx context.Context, runReq *cvm.ComputationRunReq) {
	ac := agent.Computation{
		ID:          runReq.Id,
		Name:        runReq.Name,
		Description: runReq.Description,
		Algorithm: agent.Algorithm{
			Hash:    [32]byte(runReq.Algorithm.Hash),
			UserKey: runReq.Algorithm.UserKey,
		},
	}

	for _, ds := range runReq.Datasets {
		ac.Datasets = append(ac.Datasets, agent.Dataset{
			Hash:    [32]byte(ds.Hash),
			UserKey: ds.UserKey,
		})
	}

	for _, rc := range runReq.ResultConsumers {
		ac.ResultConsumers = append(ac.ResultConsumers, agent.ResultConsumer{
			UserKey: rc.UserKey,
		})
	}

	if err := client.svc.InitComputation(ctx, ac); err != nil {
		client.logger.Warn(err.Error())
		return
	}

	client.mu.Lock()
	defer client.mu.Unlock()

	err := client.sp.Start(ctx, agent.AgentConfig{
		Port:         runReq.AgentConfig.Port,
		Host:         runReq.AgentConfig.Host,
		CertFile:     runReq.AgentConfig.CertFile,
		KeyFile:      runReq.AgentConfig.KeyFile,
		ServerCAFile: runReq.AgentConfig.ServerCaFile,
		ClientCAFile: runReq.AgentConfig.ClientCaFile,
		AttestedTls:  runReq.AgentConfig.AttestedTls,
	}, ac)
	if err != nil {
		client.logger.Warn(err.Error())
	}

	runRes := &cvm.ClientStreamMessage_RunRes{
		RunRes: &cvm.RunResponse{
			ComputationId: runReq.Id,
			Error:         err.Error(),
		},
	}
	client.sendMessage(&cvm.ClientStreamMessage{Message: runRes})
}

func (client *CVMClient) handleStopComputation(ctx context.Context, mes *cvm.ServerStreamMessage_StopComputation) {
	msg := &cvm.ClientStreamMessage_StopComputationRes{
		StopComputationRes: &cvm.StopComputationResponse{
			ComputationId: mes.StopComputation.ComputationId,
		},
	}
	if err := client.svc.StopComputation(ctx); err != nil {
		msg.StopComputationRes.Message = err.Error()
	}

	client.mu.Lock()
	defer client.mu.Unlock()

	if err := client.sp.Stop(); err != nil {
		msg.StopComputationRes.Message = err.Error()
	}

	client.sendMessage(&cvm.ClientStreamMessage{Message: msg})
}

func (client *CVMClient) handleOutgoingMessages(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case mes := <-client.messageQueue:
			if err := client.stream.Send(mes); err != nil {
				return err
			}
		}
	}
}

func (client *CVMClient) sendMessage(mes *cvm.ClientStreamMessage) {
	ctx, cancel := context.WithTimeout(context.Background(), sendTimeout)
	defer cancel()

	select {
	case client.messageQueue <- mes:
	case <-ctx.Done():
		client.logger.Warn("Failed to send message: timeout exceeded")
	}
}

type runRequestManager struct {
	requests map[string]*runRequest
	mu       sync.Mutex
}

type runRequest struct {
	buffer    []byte
	lastChunk time.Time
	timer     *time.Timer
}

func newRunRequestManager() *runRequestManager {
	return &runRequestManager{
		requests: make(map[string]*runRequest),
	}
}

func (m *runRequestManager) addChunk(id string, chunk []byte, isLast bool) ([]byte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	req, exists := m.requests[id]
	if !exists {
		req = &runRequest{
			buffer:    make([]byte, 0),
			lastChunk: time.Now(),
			timer:     time.AfterFunc(runReqTimeout, func() { m.timeoutRequest(id) }),
		}
		m.requests[id] = req
	}

	req.buffer = append(req.buffer, chunk...)
	req.lastChunk = time.Now()
	req.timer.Reset(runReqTimeout)

	if isLast {
		delete(m.requests, id)
		req.timer.Stop()
		return req.buffer, true
	}

	return nil, false
}

func (m *runRequestManager) timeoutRequest(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.requests, id)
	// Log timeout or handle it as needed
}
