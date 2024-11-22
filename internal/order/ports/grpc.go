package ports

import (
	"context"

	"github.com/Mitsui515/gorder/common/genproto/orderpb"
	"github.com/Mitsui515/gorder/order/app"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*emptypb.Empty, error) {
	// TODO: Implement
	panic("not implemented")
}
func (G GRPCServer) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	// TODO: Implement
	panic("not implemented")
}
func (G GRPCServer) UpdateOrder(ctx context.Context, request *orderpb.Order) (*emptypb.Empty, error) {
	// TODO: Implement
	panic("not implemented")
}
