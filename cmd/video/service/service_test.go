package service

import (
	"context"
	"douyin-micro/cmd/video/dal"
	"os"

	"douyin-micro/cmd/video/minio"
	"douyin-micro/kitex_gen/video"
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
	req := &video.FeedRequset{
		LatestTime: nil,
		ReqUserId:  nil,
	}
	dal.Init()
	videos, next_time, err := NewFeedService(context.Background()).FeedService(req)
	if err != nil {
		klog.Error(err)
	}
	for _, v := range videos {
		klog.Info("Videos: ", v)
	}
	klog.Info("Next Time: ", next_time)
}
