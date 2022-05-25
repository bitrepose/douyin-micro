package db

import (
	"context"
	"douyin-micro/pkg/constants"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId        int64  `json:"user_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	Title         string `json:"title"`
	CreatedAt     int64  `json:"created_at"` // Override Model CreatedAt
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

func FeedVideo(ctx context.Context, limit int64, latest_time *int64) ([]*Video, int64, error) {
	var (
		total     int64
		res       []*Video
		conn      *gorm.DB
		next_time int64
	)

	// 处理初次刷新和常规刷新情况
	if latest_time == nil {
		cur_time := int64(time.Now().Unix())
		latest_time = &cur_time
		conn = DB.WithContext(ctx).Model(&Video{}).Where("created_at <= ?", *latest_time)
	} else {
		conn = DB.WithContext(ctx).Model(&Video{}).Where("created_at >= ?", *latest_time)
	}

	if err := conn.Count(&total).Error; err != nil {
		return res, *latest_time, err
	}

	// 防止limit溢出
	if limit > total {
		limit = total
	}

	// 按create_time降序取出video，即将投稿时间倒序播出
	if err := conn.Limit(int(limit)).Order("created_at desc").Find(&res).Error; err != nil {
		return res, *latest_time, err
	}

	next_time = res[len(res)-1].CreatedAt // 将video_list中最早的视频投稿时间作为next_time

	return res, next_time, nil
}
