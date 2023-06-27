package data

import (
	"context"

	"shopping/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Create(ctx context.Context, info *biz.UserInfo) (*struct{}, error) {
	query := "INSERT INTO `user` (`name`, `password`) VALUES (?, ?)"
	_, err := r.data.db.ExecContext(ctx, query, info.UserName, info.Password)
	if err != nil {
		r.log.Errorf("Create user err:%v", err)
		return nil, err
	}
	return &struct{}{}, nil
}

func (r *userRepo) Get(ctx context.Context, info *biz.UserInfo) (*biz.User, error) {
	query := "SELECT `id`, `name` FROM `user` WHERE `name` = ? AND `password` = ?"
	res := r.data.db.QueryRowContext(ctx, query, info.UserName, info.Password)

	var data biz.User
	err := res.Scan(&data.Id, &data.Name)
	if err != nil {
		r.log.Errorf("Get user scan err:%v", err)
		return nil, err
	}
	return &data, nil
}

func (r *userRepo) GetById(ctx context.Context, id int) (*biz.User, error) {
	query := "SELECT `id`, `name` FROM `user` WHERE `id` = ?"
	res := r.data.db.QueryRowContext(ctx, query, id)

	var data biz.User
	err := res.Scan(&data.Id, &data.Name)
	if err != nil {
		r.log.Errorf("Get user scan err:%v", err)
		return nil, err
	}
	return &data, nil
}
