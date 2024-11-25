package service

import (
	"context"

	"github.com/Mitsui515/gorder/common/metrics"
	"github.com/Mitsui515/gorder/order/adapters"
	"github.com/Mitsui515/gorder/order/app"
	"github.com/Mitsui515/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricsClient),
		},
	}
}
