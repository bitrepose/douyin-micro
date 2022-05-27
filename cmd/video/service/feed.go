package service

import (
	"context"
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/cmd/video/pack"
	"douyin-micro/cmd/video/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/kitex_gen/video"
)

const (
	LIMIT = 30 // 单次返回最大视频数
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) FeedService(req *video.FeedRequset) ([]*video.Video, *int64, error) {
	videoModels, next_time, err := db.FeedVideo(s.ctx, LIMIT, req.LatestTime)
	if err != nil {
		return nil, nil, err
	}
	uIds := pack.UserIds(videoModels)
	userMap, err := rpc.MUserInfo(s.ctx, &user.MUserInfoRequest{
		UserIds: uIds,
	})
	if err != nil {
		return nil, nil, err
	}
	videos := pack.Videos(videoModels, userMap)
	if req.ReqUserId != nil {
		favVideos, err := db.FavoriteIdList(s.ctx, *req.ReqUserId)
		if err != nil {
			return nil, nil, err
		}
		for _, v := range videos {
			// 已点赞
			if _, ok := favVideos[v.Id]; ok == true {
				v.IsFavorite = true
			}
		}
	}

	return videos, &next_time, nil
}
