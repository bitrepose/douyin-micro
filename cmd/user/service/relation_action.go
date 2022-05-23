package service

import (
	"context"
	"douyin-micro/cmd/user/dal/db"
	"douyin-micro/kitex_gen/user"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *user.RelationActionRequest) error {
	actionType := req.ActionType
	userId := req.UserId
	toUserId := req.ToUserId
	var err error
	if actionType == 1 {
		err = db.Follow(s.ctx, userId, toUserId)
	} else {
		err = db.DisFollow(s.ctx, userId, toUserId)
	}
	return err
}
