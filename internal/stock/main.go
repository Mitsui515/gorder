package main

import (
	"github.com/Mitsui515/gorder/common/genproto/stockpb"
	"github.com/Mitsui515/gorder/common/server"
	"github.com/Mitsui515/gorder/stock/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer()
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// TODO: Implement
	default:
		panic("unexpected server type")
	}
}
