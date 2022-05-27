package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishList(c *gin.Context) {
	var params struct {
		UserId int64  `json:"user_id" form:"user_id"`
		Token  string `json:"token" form:"token"`
	}
	if err := c.BindQuery(&params); err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	if params.UserId < 0 {
		sendBaseResp(c, errno.ParamErr)
	}
	req := video.PublishListRequest{
		UserId: params.UserId,
	}
	resp, err := rpc.PublishList(context.Background(), &req)
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}

	c.JSON(http.StatusOK, resp)
}
