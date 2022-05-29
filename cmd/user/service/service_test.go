package service

import (
	"context"
	"douyin-micro/cmd/user/dal"
	"douyin-micro/kitex_gen/user"
	"fmt"
	"testing"
)

func TestUserRegister(t *testing.T) {
	req := &user.UserRegisterRequest{
		Username: "ylh",
		Password: "123",
	}
	dal.Init()
	token, err := NewUserRegisterService(context.Background()).UserRegister(req)
	fmt.Println(token, err)
}

func TestUserLogin(t *testing.T) {
	req := &user.UserLoginRequest{
		Username: "ylh123",
		Password: "123456",
	}
	dal.Init()
	_, err := NewUserLoginService(context.Background()).UserLogin(req)
	// fmt.Println(id)
	fmt.Println(err)
}
