// Code generated by Kitex v0.3.1. DO NOT EDIT.

package videoservice

import (
	"context"
	"douyin-micro/kitex_gen/video"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":           kitex.NewMethodInfo(feedHandler, newVideoServiceFeedArgs, newVideoServiceFeedResult, false),
		"PublishAction":  kitex.NewMethodInfo(publishActionHandler, newVideoServicePublishActionArgs, newVideoServicePublishActionResult, false),
		"PublishList":    kitex.NewMethodInfo(publishListHandler, newVideoServicePublishListArgs, newVideoServicePublishListResult, false),
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newVideoServiceFavoriteActionArgs, newVideoServiceFavoriteActionResult, false),
		"FavoriteList":   kitex.NewMethodInfo(favoriteListHandler, newVideoServiceFavoriteListArgs, newVideoServiceFavoriteListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFeedArgs)
	realResult := result.(*video.VideoServiceFeedResult)
	success, err := handler.(video.VideoService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedArgs() interface{} {
	return video.NewVideoServiceFeedArgs()
}

func newVideoServiceFeedResult() interface{} {
	return video.NewVideoServiceFeedResult()
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishActionArgs)
	realResult := result.(*video.VideoServicePublishActionResult)
	success, err := handler.(video.VideoService).PublishAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishActionArgs() interface{} {
	return video.NewVideoServicePublishActionArgs()
}

func newVideoServicePublishActionResult() interface{} {
	return video.NewVideoServicePublishActionResult()
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishListArgs)
	realResult := result.(*video.VideoServicePublishListResult)
	success, err := handler.(video.VideoService).PublishList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishListArgs() interface{} {
	return video.NewVideoServicePublishListArgs()
}

func newVideoServicePublishListResult() interface{} {
	return video.NewVideoServicePublishListResult()
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFavoriteActionArgs)
	realResult := result.(*video.VideoServiceFavoriteActionResult)
	success, err := handler.(video.VideoService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFavoriteActionArgs() interface{} {
	return video.NewVideoServiceFavoriteActionArgs()
}

func newVideoServiceFavoriteActionResult() interface{} {
	return video.NewVideoServiceFavoriteActionResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFavoriteListArgs)
	realResult := result.(*video.VideoServiceFavoriteListResult)
	success, err := handler.(video.VideoService).FavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFavoriteListArgs() interface{} {
	return video.NewVideoServiceFavoriteListArgs()
}

func newVideoServiceFavoriteListResult() interface{} {
	return video.NewVideoServiceFavoriteListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, req *video.FeedRequset) (r *video.FeedResponse, err error) {
	var _args video.VideoServiceFeedArgs
	_args.Req = req
	var _result video.VideoServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishAction(ctx context.Context, req *video.PublishActionRequest) (r *video.PublishActionResponse, err error) {
	var _args video.VideoServicePublishActionArgs
	_args.Req = req
	var _result video.VideoServicePublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, req *video.PublishListRequest) (r *video.PublishListResponse, err error) {
	var _args video.VideoServicePublishListArgs
	_args.Req = req
	var _result video.VideoServicePublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteAction(ctx context.Context, req *video.FavoriteActionRequest) (r *video.FavoriteActionResponse, err error) {
	var _args video.VideoServiceFavoriteActionArgs
	_args.Req = req
	var _result video.VideoServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, req *video.FavoriteListRequest) (r *video.FavoriteListResponse, err error) {
	var _args video.VideoServiceFavoriteListArgs
	_args.Req = req
	var _result video.VideoServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}