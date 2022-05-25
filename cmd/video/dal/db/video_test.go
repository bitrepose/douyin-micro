package db

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestFeedVideo(t *testing.T) {
	Init()
	videos, next_time, err := FeedVideo(context.Background(), 30, nil)
	if err != nil {
		klog.Error(err)
	}
	for _, v := range videos {
		klog.Info("Video Publish Time: ", v.CreatedAt)
	}
	klog.Info("Next Time: ", next_time)

	latest_time := int64(1653304653)
	videos, next_time, err = FeedVideo(context.Background(), 30, &latest_time)
	if err != nil {
		klog.Error(err)
	}
	for _, v := range videos {
		klog.Info("Video Publish Time: ", v.CreatedAt)
	}
	klog.Info("Next Time: ", next_time)

}
