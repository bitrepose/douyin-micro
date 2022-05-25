package service

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"douyin-micro/pkg/constants"
)

func PostComment(ctx context.Context, cmt *db.Comment) (*db.Comment,constants.PostCommentMessage) {
		
}

func DeleteComent(ctx context.Context,cmt *db.Comment)constants.DeleteCommentMessage{

}
