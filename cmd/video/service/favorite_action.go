package service

import (
	"context"
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/kitex_gen/video"
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
	isFavorite := false
	if req.ActionType == 1 {
		isFavorite = true
	}
	favoriteModel := &db.Favorite{
		UserId:     int(req.UserId),
		VideoId:    int(req.VideoId),
		IsFavorite: isFavorite,
	}
	return db.UpdateFavorite(s.ctx, favoriteModel)
}
