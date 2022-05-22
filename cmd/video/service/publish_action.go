package service

import (
	"bytes"
	"context"
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/cmd/video/minio"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/constants"
	"strings"

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
	fileName := u2.String()
	// 上传视频
	err = minio.UploadFile(constants.MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		return err
	}
	// 获取视频链接
	url, err := minio.GetFileUrl(constants.MinioVideoBucketName, fileName, 0)
	playUrl := strings.Split(url.String(), "?")[0]
	if err != nil {
		return err
	}
	// 封装video
	videoModel := &db.Video{
		UserId:        int(req.UserId),
		PlayUrl:       playUrl,
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	return db.CreateVideo(s.ctx, videoModel)
}
