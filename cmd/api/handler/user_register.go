package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var info LoginInfo
	if err := c.ShouldBind(&info); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
		return
	}
	if len(info.Username) == 0 || len(info.Password) == 0 {
		SendBaseResp(c, errno.ParamErr)
		return
	}
	id, token, err := rpc.Register(context.Background(), info.Username, info.Password)
	sendUserResp(c, token, id, err)

}

type LoginInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// 适用于 用户登陆与用户注册的相应
func sendUserResp(c *gin.Context, token string, userId int64, err errno.ErrNo) {
	c.JSON(http.StatusOK, UserResp{
		Response{StatusCode: int(err.ErrCode), StatusMsg: err.ErrMsg},
		userId,
		token})
}

type UserResp struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
