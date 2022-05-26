package main

import (
	"context"
	"douyin-micro/cmd/video/dal/db"
	"douyin-micro/kitex_gen/comment"
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
	db.InitDB()
	s := &CommentServiceImpl{}
	text := string("bigrain")
	creq := &comment.CommentActionRequest{
		UserId:      1,
		VideoId:     1,
		ActionType:  1,
		CommentText: &text,
		CommentId:   nil,
	}
	cresp, _ := s.CommentAction(context.Background(), creq)
	fmt.Println(cresp.StatusCode, cresp.StatusMsg, cresp.Comment)
}
