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
	return &struct{}{}, nil
}

func (r *userRepo) Get(ctx context.Context, info *biz.UserInfo) (*biz.User, error) {
	user := biz.User{
		Id:   1,
		Name: "a",
	}
	return &user, nil
}
