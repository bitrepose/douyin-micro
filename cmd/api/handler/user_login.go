package handler

import (
	"context"
	jwtutil "douyin-micro/cmd/api/jwt_util"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/pkg/errno"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var info LoginInfo
	if err := c.ShouldBind(&info); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
		return
	}
	if len(info.Username) == 0 || len(info.Password) == 0 {
		SendBaseResp(c, errno.ParamErr)
		return
	}
	id, err := rpc.Login(context.Background(), info.Username, info.Password)
	if err.ErrCode != errno.SuccessCode {
		SendBaseResp(c, err)
	}
	tokenStr, e := jwtutil.CreateToken(id)
	if e != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	sendUserResp(c, tokenStr, id, err)
}
