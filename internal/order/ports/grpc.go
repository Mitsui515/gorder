package ports

import (
	"context"

	"github.com/Mitsui515/gorder/common/genproto/orderpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
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
