package db

import (
	"context"
	"douyin-micro/pkg/constants"
	"errors"
	"time"

	"gorm.io/gorm"
)

type FavoriteRelation struct {
	UserId    int64 `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	VideoId   int64 `json:"video_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (f *FavoriteRelation) TableName() string {
	return constants.FavoriteTableName
}

func Favorite(ctx context.Context, userId int64, videoId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 新增点赞关系
		e := tx.Create(&FavoriteRelation{UserId: userId, VideoId: videoId}).Error
		if e != nil {
			return e
		}
		//2.改变 video 表中的 favorite count
		res := tx.Model(new(Video)).Where("ID = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errors.New("favorite update error")
		}
		// 返回 nil 提交事务
		if e != nil {
			return e
		}
		return nil
	})
	return err
}

func DisFavorite(ctx context.Context, userId int64, videoId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 删除点赞关系
		e := tx.Delete(&FavoriteRelation{UserId: userId, VideoId: videoId}).Error
		if e != nil {
			return e
		}
		//2.改变 video 表中的 favorite count
		e = tx.Model(new(Video)).Where("ID = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if e != nil {
			return e
		}
		// 返回 nil 提交事务
		if e != nil {
			return e
		}
		return nil
	})
	return err
}

func FavoriteList(ctx context.Context, userId int64) ([]*Video, error) {
	var favList []*Video
	err := DB.WithContext(ctx).Table("video").Joins("inner join user_video as uv on video.id = uv.video_id").Where("uv.user_id = ?", userId).Find(&favList).Error
	if err != nil {
		return nil, err
	}
	return favList, nil
}

func FavoriteIdList(ctx context.Context, userId int64) (map[int64]any, error) {
	var favList []int64
	err := DB.WithContext(ctx).Raw("SELECT v.id FROM video as v inner join user_video as uv on v.id = uv.video_id WHERE uv.user_id = ?", userId).Scan(&favList).Error
	if err != nil {
		return nil, err
	}
	favSet := make(map[int64]any)
	for _, m := range favList {
		favSet[m] = struct{}{}
	}
	return favSet, nil
}
