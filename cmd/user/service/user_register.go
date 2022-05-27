package service

import (
	"context"
	"crypto/md5"
	"douyin-micro/cmd/user/dal/db"
	"douyin-micro/kitex_gen/user"
	"fmt"
	"io"
)

type UserRegisterService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

// CreateUser create user info.
func (s *UserRegisterService) UserRegister(req *user.UserRegisterRequest) (int64, error) {
	h := md5.New()
	_, err := io.WriteString(h, req.Password)
	if err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(passWord)
	err = db.CreateUser(s.ctx, &db.User{
		Name:          req.Username,
		Password:      passWord,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	})
	user, _ := db.FindUserByUsername(s.ctx, req.Username)
	return int64(user.ID), err
}
