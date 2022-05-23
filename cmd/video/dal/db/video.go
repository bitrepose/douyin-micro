package db

import (
	"context"
	"douyin-micro/pkg/constants"
	"time"

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
	CreateTime    int64  `json:"create_time" gorm:"autoCreateTime"` // 待加入的字段
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

func FeedVideo(ctx context.Context, limit, next_time int) ([]*Video, error) {
	var total int64
	var res []*Video
	var conn *gorm.DB

	// 处理初次刷新和常规刷新情况
	if next_time == 0 {
		next_time = int(time.Now().Unix())
		conn = DB.WithContext(ctx).Model(&Video{}).Where("create_time <= ?", next_time)
	} else {
		conn = DB.WithContext(ctx).Model(&Video{}).Where("create_time >= ", next_time)
	}

	if err := conn.Count(&total).Error; err != nil {
		return res, err
	}

	// prevent offset overflow
	if limit > int(total) {
		limit = int(total)
	}

	// retrieve videos in "id desc" order so that latest video comes first
	if err := conn.Limit(limit).Order("create_time desc").Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
