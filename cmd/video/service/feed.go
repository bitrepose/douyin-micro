package service

import (
	"context"
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/cmd/video/pack"
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
	// latest_time不填则标记为0,便于dao层处理
	if req.LatestTime == nil {
		*req.LatestTime = 0
	}

	videoModels, next_time, err := db.FeedVideo(s.ctx, LIMIT, *req.LatestTime)
	if err != nil {
		return nil, nil, err
	}

	videos := pack.Videos(videoModels)

	return videos, &next_time, nil
}
