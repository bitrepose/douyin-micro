package db

import (
	"context"
	"douyin-micro/pkg/constants"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId        int    `json:"user_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	Title         string `json:"title"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}
