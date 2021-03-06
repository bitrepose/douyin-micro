package main

import (
	"context"
	"douyin-micro/cmd/video/service"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequset) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FeedResponse)

	if (req.ReqUserId != nil && *req.ReqUserId == 0) || (req.LatestTime != nil && *req.LatestTime < 0) {
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = &errno.ParamErr.ErrMsg
		return resp, nil
	}

	videos, next_time, err := service.NewFeedService(ctx).FeedService(req)
	if err != nil {
		errMsg := err.Error()
		resp.StatusCode = errno.ServiceErrCode
		resp.StatusMsg = &errMsg
		return resp, nil
	}

	resp.StatusCode = errno.SuccessCode
	resp.VideoList = videos
	resp.NextTime = next_time
	// req.SetLatestTime(next_time) // resp中的next_time作为下一次req的latest_time
	return resp, nil
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(video.PublishActionResponse)

	if len(req.Data) == 0 || len(req.Title) == 0 {
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = &errno.ParamErr.ErrMsg
		return resp, nil
	}

	err = service.NewPublishActionService(ctx).PublishAction(req)
	if err != nil {
		errMsg := err.Error()
		resp.StatusCode = errno.ServiceErrCode
		resp.StatusMsg = &errMsg
		return resp, nil
	}
	resp.StatusCode = errno.SuccessCode
	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.PublishListResponse)

	if req.UserId == 0 || (req.ReqUserId != nil && *req.ReqUserId == 0) {
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = &errno.ParamErr.ErrMsg
	}
	videos, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		errMsg := err.Error()
		resp.StatusCode = errno.ServiceErrCode
		resp.StatusMsg = &errMsg
		return resp, nil
	}
	resp.StatusCode = errno.SuccessCode
	resp.VideoList = videos

	return resp, nil
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *video.FavoriteActionRequest) (resp *video.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	if req.ActionType != 1 && req.ActionType != 2 || req.UserId == 0 || req.VideoId == 0 {
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = &errno.ParamErr.ErrMsg
	}
	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		errMsg := err.Error()
		resp.StatusCode = errno.ServiceErrCode
		resp.StatusMsg = &errMsg
		return resp, nil
	}
	resp.StatusCode = errno.SuccessCode
	return resp, nil
}

// FavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteList(ctx context.Context, req *video.FavoriteListRequest) (resp *video.FavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FavoriteListResponse)

	if req.UserId == 0 || req.ReqUserId == 0 {
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = &errno.ParamErr.ErrMsg
	}
	videos, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		errMsg := err.Error()
		resp.StatusCode = errno.ServiceErrCode
		resp.StatusMsg = &errMsg
		return resp, nil
	}
	resp.StatusCode = errno.SuccessCode
	resp.VideoList = videos
	return resp, nil
}
