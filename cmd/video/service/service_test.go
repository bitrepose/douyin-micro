package service

import (
	"context"
	"douyin-micro/cmd/video/dal"
	"douyin-micro/cmd/video/minio"
	"douyin-micro/cmd/video/rpc"
	"douyin-micro/kitex_gen/video"
	"os"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestPublishAction(t *testing.T) {
	f, _ := os.Open("../minio/test.mp4")
	defer f.Close()
	fi, _ := os.Stat("../minio/test.mp4")
	data := make([]byte, fi.Size())
	n, _ := f.Read(data)
	klog.Info(n)
	reqData := make([]int8, fi.Size())
	// uint8 -> int8
	for i, d := range data {
		reqData[i] = int8(d)
	}
	req := &video.PublishActionRequest{
		UserId: 1,
		Title:  "zjs",
		Data:   reqData,
	}
	// klog.Info(req)
	minio.InitMinio()
	dal.Init()
	err := NewPublishActionService(context.Background()).PublishAction(req)
	klog.Error(err)
}

func TestFeedService(t *testing.T) {
	req := &video.FeedRequset{}
	dal.Init()
	rpc.InitRPC()
	videos, next_time, err := NewFeedService(context.Background()).FeedService(req)
	if err != nil {
		klog.Error(err)
	}
	for _, v := range videos {
		klog.Info("Videos: ", v)
	}
	klog.Info("Next Time: ", next_time)
}

func TestPublishList(t *testing.T) {
	ReqId := int64(2)
	req := &video.PublishListRequest{
		UserId:    1,
		ReqUserId: &ReqId,
	}
	dal.Init()
	rpc.InitRPC()
	videos, err := NewPublishListService(context.Background()).PublishList(req)
	if err != nil {
		klog.Error(err)
	}
	for _, v := range videos {
		klog.Info("Videos ", v)
		klog.Info("Video Author: ", v.Author)
		klog.Info("Videos is_favorite: ", v.IsFavorite)
	}
}
