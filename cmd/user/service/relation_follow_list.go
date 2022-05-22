package service

import (
	"context"
	"douyin-micro/cmd/user/dal/db"
	"douyin-micro/cmd/user/pack"
	"douyin-micro/kitex_gen/user"
)

type RelationFollowListService struct {
	ctx context.Context
}

func NewRelationFollowListService(ctx context.Context) *RelationFollowListService {
	return &RelationFollowListService{ctx: ctx}
}

func (s *RelationFollowListService) RelationFollowList(req *user.RelationFollowListRequest) ([]*user.User, error) {
	followedIds, err := db.GetFollowedId(s.ctx, req.UserId)
	var users = []*user.User{}
	if err != nil {
		// 这里如果有异常 且类型为ErrNotFoundRecord 则表明这个用户没有关注任何人
		return users, nil
	}
	usersResult, err2 := db.FindUsersByIds(s.ctx, followedIds)
	if err2 != nil {
		// 这里有异常 表明从user表中查找用户出错
		return nil, err2
	}
	for _, tempUser := range *usersResult {
		tempUser.IsFollow = true
		users = append(users, pack.ConvUser(&tempUser))
	}
	return users, nil
}
