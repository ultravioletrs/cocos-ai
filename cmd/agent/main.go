// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	mglog "github.com/absmach/magistrala/logger"
	"github.com/absmach/magistrala/pkg/prometheus"
	"github.com/caarlos0/env/v11"
	"github.com/google/go-sev-guest/client"
	"github.com/ultravioletrs/cocos/agent"
	"github.com/ultravioletrs/cocos/agent/api"
	"github.com/ultravioletrs/cocos/agent/cvms"
	cvmapi "github.com/ultravioletrs/cocos/agent/cvms/api/grpc"
	"github.com/ultravioletrs/cocos/agent/cvms/server"
	"github.com/ultravioletrs/cocos/agent/events"
	agentlogger "github.com/ultravioletrs/cocos/internal/logger"
	"github.com/ultravioletrs/cocos/pkg/attestation/quoteprovider"
	pkggrpc "github.com/ultravioletrs/cocos/pkg/clients/grpc"
	cvmgrpc "github.com/ultravioletrs/cocos/pkg/clients/grpc/cvm"
	"golang.org/x/sync/errgroup"
)

const (
	svcName          = "agent"
	defSvcGRPCPort   = "7002"
	retryInterval    = 5 * time.Second
	envPrefixCVMGRPC = "AGENT_CVM_GRPC_"
)

type config struct {
	LogLevel string `env:"AGENT_LOG_LEVEL" envDefault:"debug"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to load %s configuration : %s", svcName, err)
	}

	var exitCode int
	defer mglog.ExitWithError(&exitCode)

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to load %s configuration : %s", svcName, err)
	}

	var level slog.Level
	if err := level.UnmarshalText([]byte(cfg.LogLevel)); err != nil {
		log.Println(err)
		exitCode = 1
		return
	}

	eventsLogsQueue := make(chan *cvms.ClientStreamMessage, 1000)

	handler := agentlogger.NewProtoHandler(os.Stdout, &slog.HandlerOptions{Level: level}, eventsLogsQueue)
	logger := slog.New(handler)

	eventSvc, err := events.New(svcName, eventsLogsQueue)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create events service %s", err.Error()))
		exitCode = 1
		return
	}

	qp, err := quoteprovider.GetQuoteProvider()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create quote provider %s", err.Error()))
		exitCode = 1
		return
	}

	cvmGrpcConfig := pkggrpc.CVMClientConfig{}
	if err := env.ParseWithOptions(&cvmGrpcConfig, env.Options{Prefix: envPrefixCVMGRPC}); err != nil {
		logger.Error(fmt.Sprintf("failed to load %s gRPC client configuration : %s", svcName, err))
		exitCode = 1
		return
	}

	cvmGRPCClient, cvmClient, err := cvmgrpc.NewCVMClient(cvmGrpcConfig)
	if err != nil {
		logger.Error(err.Error())
		exitCode = 1
		return
	}
	defer cvmGRPCClient.Close()

	pc, err := cvmClient.Process(ctx)
	if err != nil {
		logger.Error(err.Error())
		exitCode = 1
		return
	}

	svc := newService(ctx, logger, eventSvc, qp)

	mc := cvmapi.NewClient(pc, svc, eventsLogsQueue, logger, server.NewServer(logger, svc))

	g.Go(func() error {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(ch)

		select {
		case <-ch:
			logger.Info("Received signal, shutting down...")
			cancel()
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	g.Go(func() error {
		return mc.Process(ctx, cancel)
	})

	if err := g.Wait(); err != nil {
		logger.Error(fmt.Sprintf("%s service terminated: %s", svcName, err))
	}
}

func newService(ctx context.Context, logger *slog.Logger, eventSvc events.Service, qp client.QuoteProvider) agent.Service {
	svc := agent.New(ctx, logger, eventSvc, qp)

	svc = api.LoggingMiddleware(svc, logger)
	counter, latency := prometheus.MakeMetrics(svcName, "api")
	svc = api.MetricsMiddleware(svc, counter, latency)

	return svc
}
