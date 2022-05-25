package dal

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"fmt"
	"testing"
)

func TestTemp(t *testing.T) {
	Init()
	// _,err:=db.FindCommentByCommentId(context.Background(),1)
	flag:=db.DeleteCommentByCommentId(context.Background(),1)
	fmt.Printf("flag: %v\n", flag)
}
