package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Order 结构是订单模型
type Order struct {
	OrderId   int
	UserId    int
	OrderName string
}

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// OrderRepo 接口是工单仓库接口
type OrderRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// OrderUseCase is a Greeter use case.
type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

// NewOrderUseCase new a Greeter use case.
func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *OrderUseCase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
