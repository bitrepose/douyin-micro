package service

import (
	"bytes"
	"context"
	"douyin-micro/cmd/video/minio"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/constants"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gofrs/uuid"
)

type PublishActionService struct {
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{
		ctx: ctx,
	}
}

func (s *PublishActionService) PublishAction(req *video.PublishActionRequest) error {
	videoData := make([]byte, len(req.Data))
	// int8 -> uint8
	for i, d := range req.Data {
		videoData[i] = byte(d)
	}
	// byte[] -> reader
	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()
	if err != nil {
		klog.Error("uuid generate failed,", err)
		return err
	}
	err = minio.UploadFile(constants.MinioVideoBucketName, u2.String(), reader, int64(len(videoData)))
	return err
}
