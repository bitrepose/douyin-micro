package main

import (
	"context"
	"douyin-micro/cmd/user/dal"
	"douyin-micro/kitex_gen/user"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestRegister(t *testing.T) {
	dal.Init()
	service := new(UserServiceImpl)
	resp, err := service.UserRegister(context.Background(), &user.UserRegisterRequest{
		Username: "neo",
		Password: "123456",
	})
	klog.Info(resp, err)
}

func TestUserInfo(t *testing.T) {
	dal.Init()
	service := new(UserServiceImpl)
	// var reqUid int64 = 34
	resp, err := service.MUserInfo(context.Background(), &user.MUserInfoRequest{
		UserIds: []int64{1, 3},
		// ReqUserId: &reqUid,
	})
	klog.Info(resp, err)
}
