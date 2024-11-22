package ports

import (
	"context"

	"github.com/Mitsui515/gorder/common/genproto/stockpb"
)

type GRPCServer struct {
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	// TODO: Implement
	panic("not implemented")
}

func (G GRPCServer) CheckIfItemInStock(ctx context.Context, request *stockpb.CheckIfItemInStockRequest) (*stockpb.CheckIfItemInStockResponse, error) {
	// TODO: Implement
	panic("not implemented")
}
