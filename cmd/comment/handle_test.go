package main

import (
	"context"
	"douyin-micro/cmd/comment/dal"
	"douyin-micro/kitex_gen/comment"
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
	/*
		finish to test
	*/
	dal.Init()
	s := &CommentServiceImpl{}
	videos:= []int64{2}
	creq := &comment.MCommentNumberRequset{
		VideoIds:videos,
	}
	cresp, _ := s.MCommentNumber(context.Background(), creq)
	fmt.Printf("%v\n",cresp)
	fmt.Println(cresp.StatusCode, cresp.StatusMsg)
	for _,val:=range cresp.CommentNumbers{
		fmt.Println(val)
	}
}
