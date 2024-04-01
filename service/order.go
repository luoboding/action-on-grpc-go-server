package service

import (
	orders "action-on-grpc-server/orders"
	"context"
)

type OrderService struct {
	orders.UnimplementedOrderServiceServer
}

func (os *OrderService) CreateOrder(ctx context.Context, in *orders.OrderCreateRequest) (*orders.OrderResponse, error) {
	return &orders.OrderResponse{
		OrderNo: "sn-011-1400-2300-xa101212",
		UserId:  12,
	}, nil
}

func (os *OrderService) ListOrders(ctx context.Context, in *orders.ListOrdersRequest) (*orders.OrdersResponse, error) {
	return &orders.OrdersResponse{}, nil
}

func (os *OrderService) OrderDetail(ctx context.Context, in *orders.OrderDetailRequest) (*orders.OrderDetailResponse, error) {
	return &orders.OrderDetailResponse{}, nil
}
