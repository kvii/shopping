package data

import (
	"context"

	"shopping/app/order/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *orderRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *orderRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *orderRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *orderRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *orderRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
