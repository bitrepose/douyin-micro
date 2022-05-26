package service

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"douyin-micro/pkg/constants"
	"fmt"
)

func PostComment(ctx context.Context, cmt *db.Comment) constants.PostCommentMessage {
	fmt.Printf("\"1\": %v\n", "1")
	err := db.CreateComment(ctx, cmt)
	fmt.Printf("\"2\": %v\n", "2")
	if err != nil {
		return constants.PostCommentFailure
	}
	err = db.ADDCommentNumberByVideoId(ctx,int64(cmt.VideoId))
	if err != nil {
		return constants.PostCommentFailure
	}
	return constants.PostCommentSuccess
}

func DeleteComent(ctx context.Context, cmt int) constants.DeleteCommentMessage {
	success := db.DeleteCommentByCommentId(ctx, cmt)
	if success {
		return constants.DeleteCommentSuccess
	} else {
		return constants.DeleteCommentFailure
	}
}
func GetCommentList(ctx context.Context, videoId int) ([]db.Comment, constants.GetCommentListMessage) {
	comments, err := db.FindCommentSByVideoId(ctx, videoId)
	if err != nil {
		return nil ,constants.GetCommentListFailure
	}
	return comments ,constants.GetCommentListSuccess
}
