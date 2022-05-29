package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FavoriteList(c *gin.Context) {
	var params struct {
		UserId int64  `json:"user_id" form:"user_id"`
		Token  string `json:"token" form:"token"`
	}
	if err := c.BindQuery(&params); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	if params.UserId < 0 {
		SendBaseResp(c, errno.ParamErr)
	}
	req := video.FavoriteListRequest{
		UserId:    params.UserId,
		ReqUserId: c.GetInt64("uid"),
	}
	resp, err := rpc.FavoriteList(context.Background(), &req)
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}

	c.JSON(http.StatusOK, resp)
}
