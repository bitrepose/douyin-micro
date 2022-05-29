package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var info UserInfoRequest
	if err := c.ShouldBind(&info); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
		return
	}
	if info.UserId <= 0 {
		SendBaseResp(c, errno.ParamErr)
		return
	}
	//token 中获取
	userId := c.GetInt64("uid")
	klog.Info(userId)
	u, err := rpc.UserInfo(context.Background(), userId, info.UserId)
	c.JSON(http.StatusOK, InfoResponse{
		Response{StatusCode: int(err.ErrCode), StatusMsg: err.ErrMsg},
		u,
	})
}

type UserInfoRequest struct {
	UserId int64 `json:"user_id" form:"user_id"`
}

type InfoResponse struct {
	Response
	U user.User `json:"user"`
}
