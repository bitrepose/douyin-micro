package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/cmd/user/dal/db"
)

type UserRegisterService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

// CreateUser create user info.
func (s *UserRegisterService) UserRegister(req *user.UserRegisterRequest) error {
	h := md5.New()
	_, err := io.WriteString(h, req.Password)
	if err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(passWord)
	return db.CreateUser(s.ctx, &db.User{
		Name: req.Username,
		Password: passWord,
		FollowCount: 0,
		FollowerCount: 0,
		IsFollow: false,
	})
}
