package service

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"testing"
)

func TestPostComment(t *testing.T) {
	ct := &db.Comment{
		VideoId: 2,
		Text:    "nice",
		UserId:  3,
	}
	// fmt.Printf("ct: %v\n", ct)
	// code := PostComment(context.Background(), ct)
	// fmt.Printf("code: %v\n", code)
	db.CreateComment(context.Background(), ct)
}
