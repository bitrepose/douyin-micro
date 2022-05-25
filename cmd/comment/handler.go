package main

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"douyin-micro/cmd/comment/service"
	"douyin-micro/kitex_gen/comment"
	"douyin-micro/kitex_gen/user"
	"time"
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
		msg := code.String()
		tempUser := &user.User{
			Id:   1,
			Name: "ayu",
		}
		resp := &comment.CommentActionResponse{
			StatusCode: int32(code),
			StatusMsg:  &msg,
			Comment: &comment.Comment{
				Id:         int64(tempComment.ID),
				User:       tempUser,
				Content:    tempComment.Text,
				CreateDate: tempComment.CreatedAt.Format(time.RFC3339Nano),
			},
		}
		return resp, nil

	} else {
		
	}
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {

	return
}

// MCommentNumber implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) MCommentNumber(ctx context.Context, req *comment.MCommentNumberRequset) (resp *comment.MCommentNumberResponse, err error) {
	// TODO: Your code here...
	return
}
