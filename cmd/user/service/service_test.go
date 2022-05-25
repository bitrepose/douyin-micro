package service

import (
	"context"
	"douyin-micro/cmd/user/dal"
	"douyin-micro/kitex_gen/user"
	"testing"
	"fmt"
)

func TestUserRegister(t *testing.T) {
	req := &user.UserRegisterRequest{
		Username:  "ylh",
		Password:  "123",
	}
	dal.Init()
	err := NewUserRegisterService(context.Background()).UserRegister(req)
	fmt.Println(err)
}

func TestUserLogin(t *testing.T) {
	req := &user.UserLoginRequest{
		Username:  "ylh123",
		Password:  "123456",
	}
	dal.Init()
	_,err := NewUserLoginService(context.Background()).UserLogin(req)
	// fmt.Println(id)
	fmt.Println(err)
}

