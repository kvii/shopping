package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// UserInfo 结构是用户信息
type UserInfo struct {
	UserName string
	Password string
}

// User 结构是用户模型
type User struct {
	Id   int
	Name string
}

// UserRepo 接口是用户存储接口
type UserRepo interface {
	Create(context.Context, *UserInfo) (*struct{}, error)
	Get(context.Context, *UserInfo) (*User, error)
	GetById(context.Context, int) (*User, error)
}

// UserUseCase 结构是用户结构使用类型
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUseCase 函数创建了一个用户结构使用类型对象
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser 方法创建了一个用户
func (uc *UserUseCase) CreateUser(ctx context.Context, u *UserInfo) (*struct{}, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.UserName)
	return uc.repo.Create(ctx, u)
}

// GetUser 方法获取用户信息
func (uc *UserUseCase) GetUser(ctx context.Context, u *UserInfo) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUser: %v", u.UserName)
	return uc.repo.Get(ctx, u)
}

// GetUserById 方法获取用户信息
func (uc *UserUseCase) GetUserById(ctx context.Context, id int) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUserById: %v", id)
	return uc.repo.GetById(ctx, id)
}
