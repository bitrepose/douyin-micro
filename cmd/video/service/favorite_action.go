package service

import (
	"context"
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/errno"
	"errors"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteActionService) FavoriteAction(req *video.FavoriteActionRequest) error {
	// 1-点赞
	if req.ActionType == 1 {
		return db.Favorite(s.ctx, req.UserId, req.VideoId)
	}
	// 2-取消点赞
	if req.ActionType == 2 {
		return db.DisFavorite(s.ctx, req.UserId, req.VideoId)
	}
	return errors.New(errno.ParamErr.ErrMsg)
}
