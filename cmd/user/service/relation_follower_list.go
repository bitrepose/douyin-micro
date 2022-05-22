package service

import (
	"context"
	"douyin-micro/cmd/user/dal/db"
	"douyin-micro/cmd/user/pack"
	"douyin-micro/kitex_gen/user"
)

type RelationFollowerListService struct {
	ctx context.Context
}

func NewRelationFollowerListService(ctx context.Context) *RelationFollowerListService {
	return &RelationFollowerListService{ctx: ctx}
}

func (s *RelationFollowerListService) RelationFollowerList(req *user.RelationFollowerListRequest) ([]*user.User, error) {
	followerIds, err := db.GetFollowersId(s.ctx, req.UserId)
	var result = []*user.User{}
	if err != nil {
		return result, nil
	}
	users, err := db.FindUsersByIds(s.ctx, followerIds)
	if err != nil {
		return nil, err
	}
	for _, tempUser := range *users {
		tempUser.IsFollow, err = db.Followed(s.ctx, req.UserId, tempUser.ID)
		//if err!=nil && !errors.As(err,&gorm.ErrRecordNotFound) {
		//	return nil,err
		//}
		// ?2. 如果有异常 就先让他为false
		if err != nil {
			tempUser.IsFollow = false
		}
		result = append(result, pack.ConvUser(&tempUser))
	}
	return result, nil
}
