package service

import (
	"context"
	"douyin-micro/cmd/video/rpc"
	"douyin-micro/kitex_gen/user"
)

func GetUserByUserId(ctx context.Context, userId int64, reqUserId int64) (*user.User, error) {

	// ttt:=user.User{
	// 	Id:1,
	// 	Name:"hello",
	// }
	// return &ttt,nil
	rpc.InitRPC()
	users := []int64{int64(userId)}
	req := &user.MUserInfoRequest{
		UserIds:   users,
		ReqUserId: &reqUserId,
	}
	userInfos, err := rpc.MUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	tuser := userInfos[userId]
	return tuser, nil
}
