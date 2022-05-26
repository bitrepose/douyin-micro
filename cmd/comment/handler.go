package main

import (
	"context"
	"douyin-micro/cmd/comment/dal/db"
	"douyin-micro/cmd/comment/service"
	"douyin-micro/kitex_gen/comment"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/constants"
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
		var tempUser *user.User
		tempUser, err = service.GetUserByUserId(ctx, req.UserId, req.UserId)
		if err != nil {
			return nil, err
		}
		resp = &comment.CommentActionResponse{
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
		code := service.DeleteComent(ctx, int(*req.CommentId))
		msg := code.String()
		resp = &comment.CommentActionResponse{
			StatusCode: int32(code),
			StatusMsg:  &msg,
		}
		return resp, nil
	}
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp = &comment.CommentListResponse{}
	comments, code := service.GetCommentList(ctx, int(req.VideoId))
	msg := code.String()
	resp.StatusCode = int32(code)
	resp.StatusMsg = &msg
	res := make([]*comment.Comment, len(comments))
	tempuser, err := service.GetUserByUserId(ctx, req.UserId, req.UserId)
	if err != nil {
		return nil, err
	}
	if code == constants.GetCommentListSuccess {
		for idx := range comments {
			res[idx] = &comment.Comment{
				Id:         int64(comments[idx].ID),
				User:       tempuser,
				Content:    comments[idx].Text,
				CreateDate: comments[idx].CreatedAt.Format(time.RFC3339Nano),
			}
		}
		resp.CommentList = res
	}
	return resp, nil
}

// MCommentNumber implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) MCommentNumber(ctx context.Context, req *comment.MCommentNumberRequset) (resp *comment.MCommentNumberResponse, err error) {
	resp = &comment.MCommentNumberResponse{}
	resp.CommentNumbers = make([]int64, len(req.VideoIds))
	for idx, val := range req.VideoIds {
		num, err := db.QueryCommentNumberByVideo(ctx, val)
		if err != nil {
			resp.CommentNumbers[idx] = 0
		} else {
			resp.CommentNumbers[idx] = num
		}
	}
	return resp, nil
}
