package service

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"douyin-micro/pkg/constants"
)

func PostComment(ctx context.Context, cmt *db.Comment) (constants.PostCommentMessage) {
	err := db.CreateComment(ctx,cmt)
	if err != nil{
		return constants.PostCommentFailure
	}
	return constants.PostCommentSuccess
}

func DeleteComent(ctx context.Context,cmt int)constants.DeleteCommentMessage{
	success:=db.DeleteCommentByCommentId(ctx,cmt)
	if success {
		return constants.DeleteCommentSuccess
	}else{
		return constants.DeleteCommentFailure
	}
}
