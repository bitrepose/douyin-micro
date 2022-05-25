package db

import (
	"context"
	"douyin-micro/pkg/constants"

	"gorm.io/gorm"
)

type Comment struct {
	*gorm.Model
	VideoId int    `json:"video_id"`
	Text    string `json:"text"`
	UserId  int    `json:"user_id"`
}

func (cmt *Comment) TableName() string {
	return constants.CommentTableName
}

/*
	创建一个评论，即把一个Comment{}插入数据库
*/
func CreateComment(ctx context.Context, comment *Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}
func FindCommentSByVideoId(ctx context.Context, videoId int) ([] Comment, error) {
	var comments = []Comment{}
	result := DB.WithContext(ctx).Where("video_id = ?", videoId).Find(&comments)
	return comments, result.Error
}
