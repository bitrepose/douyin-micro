package db

import (
	"context"
	"douyin-micro/pkg/constants"
	"errors"
	"gorm.io/gorm"
)

// User 用户的模型
type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow" gorm:"-"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser 在user表中新增一个用户 参数user的id不得填写 如果用户名重复 会返回一个错误
func CreateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// FindUserByUsername 根据用户名查询用户信息 如果不存在 会返回一个错误ErrRecordNotFound
func FindUserByUsername(ctx context.Context, username string) (*User, error) {
	var res = &User{}
	result := DB.WithContext(ctx).Where("name = ?", username).First(&res)
	return res, result.Error
}

// FindUserById 根据用户id查询用户信息 如果不存在 会返回一个错误ErrRecordNotFound
func FindUserById(ctx context.Context, id int) (*User, error) {
	var res = &User{}
	result := DB.WithContext(ctx).First(&res, id)
	return res, result.Error
}

// Follow 关注操作
// @Params userId 当前用户id
// @Params toUserId 被关注用户的id
// 将两个用户的关系新增到关系表（follow_follower） 并且更新user表中对应的follow_count follower_count 如果事务没有成功commit 会返回一个错误
func Follow(ctx context.Context, userId int64, toUserId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作
		//1. 新增关注关系
		e := tx.Create(&FollowRelation{UserId: toUserId, FollowerId: userId}).Error
		if e != nil {
			return e
		}
		//2.改变 user表中的count
		result := tx.Model(&User{ID: userId}).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("user表中不存在相应数据")
		}
		result = tx.Model(&User{ID: toUserId}).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("user表中不存在相应数据")
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

// DisFollow 取消关注操作
// @Params userId 当前用户id
// @Params toUserId 被取消关注用户的id
// 将两个用户的关系从关系表删除（follow_follower） 并且更新user表中对应的follow_count follower_count 如果事务没有成功commit 会返回一个错误
func DisFollow(ctx context.Context, userId int64, toUserId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 删除关注关系
		result := tx.Delete(&FollowRelation{UserId: toUserId, FollowerId: userId})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("user表中没有相应的数据")
		}
		//2.改变 user表中的count
		result = tx.Model(&User{ID: userId}).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("user表中没有相应的数据")
		}
		result = tx.Model(&User{ID: toUserId}).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("user表中没有相应的数据")
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

// FindUsersByIds 通过一组id 查询用户 返回值为用户数组
// 注意这个方法返回的user 的 is_follow字段为默认值 需要用followed方法查找
func FindUsersByIds(ctx context.Context, ids *[]int64) (*[]User, error) {
	var users = &[]User{}
	result := DB.WithContext(ctx).Where(*ids).Find(users)
	return users, result.Error
}
