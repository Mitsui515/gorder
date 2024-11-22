package ports

import (
	"context"

	"github.com/Mitsui515/gorder/common/genproto/stockpb"
	"github.com/Mitsui515/gorder/stock/app"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	// TODO: Implement
	panic("not implemented")
}

func (G GRPCServer) CheckIfItemInStock(ctx context.Context, request *stockpb.CheckIfItemInStockRequest) (*stockpb.CheckIfItemInStockResponse, error) {
	// TODO: Implement
	panic("not implemented")
}
