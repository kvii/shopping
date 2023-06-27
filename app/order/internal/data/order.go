package data

import (
	"context"

	v1 "shopping/api/user/v1"
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

func (r *orderRepo) Create(ctx context.Context, o *biz.CreateOrder) (*struct{}, error) {
	query := "INSERT INTO `order` (`user_id`, `name`) VALUES (?, ?)"
	_, err := r.data.db.ExecContext(ctx, query, o.UserId, o.OrderName)
	if err != nil {
		r.log.Errorf("CreateOrder error: %v", err)
		return nil, err
	}
	return &struct{}{}, nil
}

func (r *orderRepo) ListById(ctx context.Context, id int) ([]biz.OrderInfo, error) {
	_, err := r.data.uc.GetUserById(ctx, &v1.GetUserByIdRequest{UserId: int64(id)})
	if err != nil {
		r.log.Errorf("ListById user id error: %v", err)
		return nil, err
	}

	query := "SELECT `id`, `name` FROM `order` WHERE `user_id` = ?"

	rows, err := r.data.db.QueryContext(ctx, query, id)
	if err != nil {
		r.log.Errorf("ListById query error: %v", err)
		return nil, err
	}

	var list []biz.OrderInfo
	for rows.Next() {
		var o biz.OrderInfo
		err := rows.Scan(&o.OrderId, &o.OrderName)
		if err != nil {
			r.log.Errorf("ListById scan error: %v", err)
			return nil, err
		}
		list = append(list, o)
	}
	return list, nil
}
