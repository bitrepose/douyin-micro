package db

import (
	"context"
	"douyin-micro/pkg/constants"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type Favorite struct {
	UserId     int  `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	VideoId    int  `json:"video_id" gorm:"primaryKey;autoIncrement:false"`
	IsFavorite bool `json:"is_favorite"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (f *Favorite) TableName() string {
	return constants.FavoriteTableName
}

func UpdateFavorite(ctx context.Context, f *Favorite) error {
	newF := new(Favorite)
	err := DB.Where("user_id = ? AND video_id = ?", f.UserId, f.VideoId).Take(newF).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return DB.WithContext(ctx).Create(f).Error
		}
		klog.Error("db went wrong,", err)
		return err
	}
	return DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", f.UserId, f.VideoId).Update("is_favorite", f.IsFavorite).Error
}
