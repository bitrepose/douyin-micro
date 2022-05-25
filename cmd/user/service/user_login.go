package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"douyin-micro/pkg/errno"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/cmd/user/dal/db"
)

type UserLoginService struct {
	ctx context.Context
}

func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{
		ctx: ctx,
	}
}

func (s *UserLoginService) UserLogin(req *user.UserLoginRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(passWord)
	userName := req.Username
	user, err := db.FindUserByUsername(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if user.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(user.ID), nil
}
