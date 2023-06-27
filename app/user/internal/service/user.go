package service

import (
	"context"

	pb "shopping/api/user/v1"
	"shopping/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user")),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	s.log.Infof("CreateUser req:%v", req)
	u := &biz.UserInfo{
		UserName: req.UserName,
		Password: req.Password,
	}
	_, err := s.uc.CreateUser(ctx, u)
	if err != nil {
		s.log.Infof("CreateUser err:%v", err)
		return nil, err
	}
	return &pb.CreateUserReply{}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	ui := &biz.UserInfo{
		UserName: req.UserName,
		Password: req.Password,
	}

	u, err := s.uc.GetUser(ctx, ui)
	if err != nil {
		s.log.Infof("GetUser err:%v", err)
		return nil, err
	}

	reply := &pb.GetUserReply{
		UserId:   int64(u.Id),
		UserName: u.Name,
	}
	return reply, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdReply, error) {
	u, err := s.uc.GetUserById(ctx, int(req.UserId))
	if err != nil {
		s.log.Infof("GetUserById err:%v", err)
		return nil, err
	}

	reply := &pb.GetUserByIdReply{
		UserId:   int64(u.Id),
		UserName: u.Name,
	}
	return reply, nil
}
