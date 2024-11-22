package main

import (
	"context"

	"github.com/Mitsui515/gorder/common/config"
	"github.com/Mitsui515/gorder/common/genproto/stockpb"
	"github.com/Mitsui515/gorder/common/server"
	"github.com/Mitsui515/gorder/stock/ports"
	"github.com/Mitsui515/gorder/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	logrus.Info(serverType)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// TODO: Implement
	default:
		panic("unexpected server type")
	}
}
