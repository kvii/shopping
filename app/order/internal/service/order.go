package service

import (
	"context"

	pb "shopping/api/order/v1"
	"shopping/app/order/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type OrderService struct {
	pb.UnimplementedOrderServer

	oc  *biz.OrderUseCase
	log *log.Helper
}

func NewOrderService(oc *biz.OrderUseCase, logger log.Logger) *OrderService {
	return &OrderService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/order")),
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	s.log.Infof("CreateOrder: %v", req)

	o := &biz.CreateOrder{
		UserId:    int(req.UserId),
		OrderName: req.OrderName,
	}
	_, err := s.oc.CreateOrder(ctx, o)
	if err != nil {
		s.log.Errorf("CreateOrder err:%v", err)
		return nil, err
	}
	return &pb.CreateOrderReply{}, nil
}

func (s *OrderService) ListOrderById(ctx context.Context, req *pb.ListOrderByIdRequest) (*pb.ListOrderByIdReply, error) {
	s.log.Infof("ListOrderById: %v", req)

	oi, err := s.oc.ListOrderById(ctx, int(req.UserId))
	if err != nil {
		s.log.Errorf("ListOrderById err:%v", err)
		return nil, err
	}

	reply := &pb.ListOrderByIdReply{
		OrderList: make([]*pb.OrderInfo, len(oi)),
	}
	for i, v := range oi {
		reply.OrderList[i] = &pb.OrderInfo{
			OrderId:   int64(v.OrderId),
			OrderName: v.OrderName,
		}
	}
	return reply, nil
}
