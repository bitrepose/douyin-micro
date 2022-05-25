package main

import (
	"context"
	"douyin-micro/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	if req.ActionType ==1 {
		
	}else{

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
