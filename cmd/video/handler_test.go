package main

import (
	"context"
	"douyin-micro/cmd/video/dal"
	"douyin-micro/cmd/video/rpc"
	"douyin-micro/kitex_gen/video"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestFeed(t *testing.T) {
	dal.Init()
	rpc.InitRPC()
	service := new(VideoServiceImpl)
	resp, err := service.Feed(context.Background(), &video.FeedRequset{})
	klog.Info(resp, err)
}
