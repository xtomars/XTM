/*
 * *******************************************************************
 * @项目名称: grpcserver
 * @文件名称: grpcserver.go
 * @Date: 2020/02/07
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package grpcserver

import (
	"net"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/config"
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/grpcserver/controller"
	"git.bhex.io/bhpc/wallet/common/logger"
	"git.bhex.io/bhpc/wallet/common/protos/chainnode"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

var (
	log        = logger.New("grpcserver")
	grpcServer *grpc.Server
)

// Run a grpc server
func Run() {
	defer stopGRPC()
	startGRPC()
}

func startGRPC() {
	// listen port
	listen, err := net.Listen(config.Config.Chainnode.Protocol, ":"+config.Config.Chainnode.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	log.Infof("baasnode server start success, listen port:%v", config.Config.Chainnode.Port)

	// Initialize gRPC server's interceptor.
	grpcServer = grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	// Register gRPC service implementations.
	chainnode.RegisterChainnodeServer(grpcServer, controller.NewBaasnode())

	// After all your registrations, make sure all of the Prometheus metrics are initialized.
	grpc_prometheus.EnableHandlingTimeHistogram(grpc_prometheus.WithHistogramBuckets(prometheus.DefBuckets))
	grpc_prometheus.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func stopGRPC() {
	grpcServer.GracefulStop()
}
