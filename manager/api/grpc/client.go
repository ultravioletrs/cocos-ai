// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/ultravioletrs/cocos/manager"
	"google.golang.org/grpc"
)

const svcName = "manager.ManagerService"

type grpcClient struct {
	run     endpoint.Endpoint
	timeout time.Duration
}

// NewClient returns new gRPC client instance.
func NewClient(conn *grpc.ClientConn, timeout time.Duration) manager.ManagerServiceClient {
	return &grpcClient{
		run: kitgrpc.NewClient(
			conn,
			svcName,
			"Run",
			encodeRunRequest,
			decodeRunResponse,
			manager.RunResponse{},
		).Endpoint(),
		timeout: timeout,
	}
}

// encodeRunRequest is a transport/grpc.EncodeRequestFunc that
// converts a user-domain runReq to a gRPC request.
func encodeRunRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(runReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type: %T", request)
	}
	return &manager.RunRequest{
		Computation: req.Computation,
		CaCerts:     req.CACerts,
		ClientTls:   req.ClientTLS,
		Timeout:     req.Timeout.String(),
	}, nil
}

// decodeRunResponse is a transport/grpc.DecodeResponseFunc that
// converts a gRPC RunResponse to a user-domain response.
func decodeRunResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	_, ok := grpcResponse.(*manager.RunResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", grpcResponse)
	}
	return runRes{}, nil
}

func nopDecoder(ctx context.Context, _ interface{}) (interface{}, error) {
	return nil, nil
}

func (client grpcClient) Run(ctx context.Context, req *manager.RunRequest, _ ...grpc.CallOption) (*manager.RunResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	dur, err := time.ParseDuration(req.GetTimeout())
	if err != nil {
		return nil, err
	}

	runReq := runReq{
		Computation: req.GetComputation(),
		ClientTLS:   req.GetClientTls(),
		CACerts:     req.GetCaCerts(),
		Timeout:     dur,
	}

	_, err = client.run(ctx, runReq)
	if err != nil {
		return nil, err
	}

	return &manager.RunResponse{}, nil
}
