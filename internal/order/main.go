package main

import (
	"context"
	"log"

	"github.com/Mitsui515/gorder/common/config"
	"github.com/Mitsui515/gorder/common/genproto/orderpb"
	"github.com/Mitsui515/gorder/common/server"
	"github.com/Mitsui515/gorder/order/ports"
	"github.com/Mitsui515/gorder/order/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)

	// 在协程中启动gRPC服务，因为grpcServer.Serve(listen)是一个阻塞的方法
	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{
			app: application,
		}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
