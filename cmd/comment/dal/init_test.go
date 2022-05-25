package dal

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"testing"
)

func TestTemp(t *testing.T) {
	Init()
	cc := &db.Comment{VideoId: 2, Text: "good", UserId: 3}
	db.CreateComment(context.Background(), cc)
	// _,err:=db.FindCommentByCommentId(context.Background(),1)
	// flag:=db.DeleteCommentByCommentId(context.Background(),1)
	// fmt.Printf("flag: %v\n", flag)
}
