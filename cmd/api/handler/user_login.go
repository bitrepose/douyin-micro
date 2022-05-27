package handler

import (
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/pkg/errno"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var info LoginInfo
	if err := c.ShouldBind(&info); err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
		return
	}
	if len(info.Username) == 0 || len(info.Password) == 0 {
		sendBaseResp(c, errno.ParamErr)
		return
	}
	id, token, err := rpc.Login(c, info.Username, info.Password)
	sendUserResp(c, token, id, err)
}
