package handler

import (
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(c *gin.Context) {
	var info UserInfoRequest
	if err := c.ShouldBind(&info); err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	if info.UserId <= 0 || len(info.Token) == 0 {
		sendBaseResp(c, errno.ParamErr)
	}
	//token 中获取
	var userId int64
	u, err := rpc.UserInfo(c, userId, info.UserId)
	c.JSON(http.StatusOK, InfoResponse{
		Response{StatusCode: int(err.ErrCode), StatusMsg: err.ErrMsg},
		u,
	})
}

type UserInfoRequest struct {
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}

type InfoResponse struct {
	Response
	U user.User `json:"user"`
}
