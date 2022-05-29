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
		UserId     int64  `json:"user_id" form:"user_id"`
		Token      string `json:"token" form:"token"`
		VideoId    int64  `json:"video_id form:"video_id"`
		ActionType int32  `json:"action_type" form:"action_type"`
	}
	if err := c.BindQuery(&params); err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	if params.UserId < 0 || params.VideoId < 0 || (params.ActionType != 1 && params.ActionType != 2) {
		sendBaseResp(c, errno.ParamErr)
	}
	req := comment.CommentActionRequest{
		UserId:     params.UserId,
		VideoId:    params.VideoId,
		ActionType: params.ActionType,
	}
	err := rpc.CommentAction(context.Background(), &req)
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	sendBaseResp(c, errno.Success)
}
