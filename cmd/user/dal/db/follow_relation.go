package db

import (
	"context"
)

//FollowRelation 关注者与被关注者关系的模型
type FollowRelation struct {
	UserId     int64 `gorm:"primaryKey" json:"user_id"`
	FollowerId int64 `json:"follower_id" gorm:"primaryKey;index:follower_index"`
}

func (f *FollowRelation) TableName() string {
	return "follow_follower"
}

// GetFollowersId 获取用户的粉丝的所有id 返回一个id数组指针 当出错时会返回一个错误
func GetFollowersId(ctx context.Context, userId int64) (*[]int64, error) {
	var res = []int64{}
	var temp = []FollowRelation{}
	result := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&temp)
	for _, relation := range temp {
		res = append(res, relation.FollowerId)
	}
	return &res, result.Error

}

// GetFollowedId 获取用户的关注者所有id 返回id数组指针
func GetFollowedId(ctx context.Context, userId int64) (*[]int64, error) {
	var res = []int64{}
	var temp = []FollowRelation{}
	result := DB.WithContext(ctx).Where("follower_id = ?", userId).Find(&temp)
	for _, relation := range temp {
		res = append(res, relation.UserId)
	}
	return &res, result.Error
}

//Followed 查询当前用户是否关注了某个用户
//@Params userId 当前用户id
//@Params toUserId 某个用户id
//return bool类型 表示当前用户是否关注了某个用户
func Followed(ctx context.Context, userId int64, toUserId int64) (bool, error) {
	res := &FollowRelation{}
	result := DB.WithContext(ctx).Where("user_id = ?", toUserId).Where("follower_id = ?", userId).First(res)
	if result.Error == nil {
		return true, nil
	}
	return false, result.Error
}
