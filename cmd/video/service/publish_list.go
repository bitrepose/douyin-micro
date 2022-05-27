package service

import (
	"context"
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/cmd/video/pack"
	"douyin-micro/cmd/video/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/kitex_gen/video"
)

type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{
		ctx: ctx,
	}
}

func (s *PublishListService) PublishList(req *video.PublishListRequest) ([]*video.Video, error) {
	videoModels, err := db.PublishList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	userMap, err := rpc.MUserInfo(s.ctx, &user.MUserInfoRequest{
		UserIds: pack.UserIds(videoModels),
	})
	videos := pack.Videos(videoModels, userMap)
	if req.ReqUserId != nil {
		favVideos, err := db.FavoriteIdList(s.ctx, *req.ReqUserId)
		if err != nil {
			return nil, err
		}
		for _, v := range videos {
			// 已点赞
			if _, ok := favVideos[v.Id]; ok == true {
				v.IsFavorite = true
			}
		}
	}
	return videos, nil
}
