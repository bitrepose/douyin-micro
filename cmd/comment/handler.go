package main

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"douyin-micro/cmd/comment/service"
	"douyin-micro/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	if req.ActionType == 1 {
		tempComment := &db.Comment{
			UserId:  int(req.UserId),
			VideoId: int(req.UserId),
			Text:    *req.CommentText,
		}
		code := service.PostComment(ctx, tempComment)
		msg:=code.String()
		myresp := &comment.CommentActionResponse{
			StatusCode: int32(code),
			StatusMsg:  &msg,
			Comment: 
		}
	} else {

	}
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {

	return
}
