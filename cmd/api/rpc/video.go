package rpc

import (
	"context"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/kitex_gen/video/videoservice"
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/errno"
	"douyin-micro/pkg/middleware"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func PublishAction(ctx context.Context, req *video.PublishActionRequest) error {
	resp, err := videoClient.PublishAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return nil
}

func PublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	resp, err := videoClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteAction(ctx context.Context, req *video.FavoriteActionRequest) error {
	resp, err := videoClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return nil
}

func FavoriteList(ctx context.Context, req *video.FavoriteListRequest) (*video.FavoriteListResponse, error) {
	resp, err := videoClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func Feed(ctx context.Context, req *video.FeedRequset) (*video.FeedResponse, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
