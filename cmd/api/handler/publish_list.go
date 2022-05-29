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
		SendBaseResp(c, errno.ConvertErr(err))
	}
	if params.UserId < 0 {
		SendBaseResp(c, errno.ParamErr)
	}
	reqUid := c.GetInt64("uid")
	req := video.PublishListRequest{
		UserId:    params.UserId,
		ReqUserId: &reqUid,
	}
	resp, err := rpc.PublishList(context.Background(), &req)
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}

	c.JSON(http.StatusOK, resp)
}
