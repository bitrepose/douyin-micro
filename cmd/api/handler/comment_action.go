package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/comment"
	"douyin-micro/pkg/errno"

	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	// TODO: optional params: comment_text , comnent_id
	var params struct {
		UserId      int64  `json:"user_id" form:"user_id"`
		VideoId     int64  `json:"video_id form:"video_id"`
		ActionType  int32  `json:"action_type" form:"action_type"`
		CommentText string `json:"comment_text" form:"comment_text"`
		CommentId   int64  `json:"comment_id" form:"comment_id"`
	}
	if err := c.BindQuery(&params); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	if params.UserId < 0 || params.VideoId < 0 || (params.ActionType != 1 && params.ActionType != 2) {
		SendBaseResp(c, errno.ParamErr)
	}
	req := comment.CommentActionRequest{
		UserId:      params.UserId,
		VideoId:     params.VideoId,
		ActionType:  params.ActionType,
		CommentText: &params.CommentText,
		CommentId:   &params.CommentId,
	}
	err := rpc.CommentAction(context.Background(), &req)
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	SendBaseResp(c, errno.Success)
}
