package service

import (
	"context"
	"douyin-micro/kitex_gen/video"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteListService) FavoriteList(req *video.FavoriteActionRequest) {

}
