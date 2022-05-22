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
	return
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
	return
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *video.FavoriteActionRequest) (resp *video.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteList(ctx context.Context, req *video.FavoriteListRequest) (resp *video.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}
