package handler

import (
	jwtutil "douyin-micro/cmd/api/jwt_util"
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
	id, err := rpc.Login(c, info.Username, info.Password)
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	tokenStr, err := jwtutil.CreateToken(id)
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	sendUserResp(c, tokenStr, id, err)
}
