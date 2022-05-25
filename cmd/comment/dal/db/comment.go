package db

import (
	"context"
	"douyin-micro/pkg/constants"
	"errors"
	"fmt"

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
/*
	根据视频id 找到所有的评论
*/
func FindCommentSByVideoId(ctx context.Context, videoId int) ([]Comment, error) {
	var comments = []Comment{}
	result := DB.WithContext(ctx).Raw("SELECT * FROM `comment` WHERE (video_id = ? and deleted_at IS NULL)",videoId).Find(&comments)
	return comments, result.Error
}
/*
	根据评论id查找某个评论
*/
func FindCommentByCommentId(ctx context.Context,commentId int)(*Comment,error){
	cmt:=& Comment{} 
	sql:=fmt.Sprintf("SELECT * FROM `comment` WHERE (id = %d and deleted_at IS NULL)",commentId )
	result:=DB.WithContext(ctx).Raw(sql).First(cmt)

	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		return nil,result.Error
	}
	return cmt,nil
}
/*
	删除某个评论
*/
func DeleteCommentByCommentId(ctx context.Context,commentId int)(bool){
	if _,err:=FindCommentByCommentId(ctx,commentId);err!=nil{
		return false
	}
	result:=DB.Delete(&Comment{},commentId)
	return result.Error == nil 
}

