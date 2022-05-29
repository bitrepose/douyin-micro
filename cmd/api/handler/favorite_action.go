package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/errno"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	var params struct {
		UserId     int64  `json:"user_id" form:"user_id"`
		Token      string `json:"token" form:"token"`
		VideoId    int64  `json:"video_id" form:"video_id"`
		ActionType int32  `json:"action_type" form:"action_type"`
	}
	if err := c.BindQuery(&params); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	if params.UserId < 0 || params.VideoId < 0 || params.ActionType != 1 && params.ActionType != 2 {
		SendBaseResp(c, errno.ParamErr)
	}
	req := video.FavoriteActionRequest{
		UserId:     params.UserId,
		VideoId:    params.VideoId,
		ActionType: params.ActionType,
	}
	err := rpc.FavoriteAction(context.Background(), &req)
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	SendBaseResp(c, errno.Success)
}
