package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// OrderInfo 结构是订单信息模型
type OrderInfo struct {
	OrderId   int
	OrderName string
}

// CreateOrder 结构是创建订单参数
type CreateOrder struct {
	UserId    int
	OrderName string
}

// OrderRepo 接口是工单仓库接口
type OrderRepo interface {
	Create(context.Context, *CreateOrder) (*struct{}, error)
	ListById(ctx context.Context, id int) ([]OrderInfo, error)
}

// OrderUseCase 结构是订单仓库用例
type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

// NewOrderUseCase 创建订单用例
func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateOrder 方法创建订单
func (uc *OrderUseCase) CreateOrder(ctx context.Context, g *CreateOrder) (*struct{}, error) {
	uc.log.WithContext(ctx).Infof("CreateOrder: %v", g)
	return uc.repo.Create(ctx, g)
}

// ListOrderById 方法获取订单信息
func (uc *OrderUseCase) ListOrderById(ctx context.Context, id int) ([]OrderInfo, error) {
	uc.log.WithContext(ctx).Infof("ListOrderById: %v", id)
	return uc.repo.ListById(ctx, id)
}
