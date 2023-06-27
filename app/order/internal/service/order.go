package service

import (
	"context"

	pb "shopping/api/order/v1"
	"shopping/app/order/internal/biz"
)

type OrderService struct {
	pb.UnimplementedOrderServer

	oc *biz.OrderUseCase
}

func NewOrderService(oc *biz.OrderUseCase) *OrderService {
	return &OrderService{
		oc: oc,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	return &pb.CreateOrderReply{}, nil
}
func (s *OrderService) ListOrderById(ctx context.Context, req *pb.ListOrderByIdRequest) (*pb.ListOrderByIdReply, error) {
	return &pb.ListOrderByIdReply{}, nil
}
