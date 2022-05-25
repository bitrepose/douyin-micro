package service

import (
	"context"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/cmd/user/dal/db"
	"douyin-micro/cmd/user/pack"
)

type MUserInfoService struct {
	ctx context.Context
}

func NewMUserInfoService(ctx context.Context) *MUserInfoService {
	return &MUserInfoService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MUserInfoService) MUserInfo(req *user.MUserInfoRequest) ([]*user.User, error) {
	modelUsers, err := db.FindUsersByIds(s.ctx, &req.UserIds)
	if err != nil {
		return nil, err
	}
	var users = []*user.User{}
	for _, tempUser := range *modelUsers {
		users = append(users, pack.ConvUser(tempUser))
	}
	return users, nil
}
