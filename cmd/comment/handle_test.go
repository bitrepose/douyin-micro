package main

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"douyin-micro/cmd/comment/service"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ct := &db.Comment{
		VideoId: 2,
		Text:    "nice",
		UserId:  3,
	}
	code := service.PostComment(context.Background(), ct)
	fmt.Printf("code: %v\n", code)
}
