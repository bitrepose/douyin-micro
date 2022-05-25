package db

import (
	"context"
	"douyin-micro/pkg/constants"
	"errors"

	"gorm.io/gorm"
)

type VideoComment struct {
	VideoId       int64 `json:"video_id" gorm:"primaryKey"`
	CommentNumber int64 `json:"comment_id" gorm:"default:0"`
}

func (c *VideoComment) TableName() string {
	return constants.VideoCommentName
}

func ADDCommentNumberByVideoId(ctx context.Context, videoId int64) error {
	vc := &VideoComment{}
	result := DB.WithContext(ctx).Where("video_id = ?", videoId).First(vc)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return DB.WithContext(ctx).Create(&VideoComment{
			VideoId: videoId,
			CommentNumber: 1,
		}).Error
	}else{
		return DB.WithContext(ctx).Model(&VideoComment{}).Where("video_id = ?",videoId).Update("comment_number",gorm.Expr("comment_number + ?",1)).Error
	}
}

func QueryCommentNumberByVideo(ctx context.Context,videoId int64)(int64,error){
	vc:= &VideoComment{}
	result:=DB.WithContext(ctx).Where("video_id = ?",videoId).First(vc)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0,nil
	}else if result.Error!=nil {
		return 0,result.Error
	}else{
		return vc.CommentNumber,nil
	}
}
