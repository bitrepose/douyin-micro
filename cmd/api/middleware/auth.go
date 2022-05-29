package middleware

import (
	"douyin-micro/cmd/api/handler"
	jwtutil "douyin-micro/cmd/api/jwt_util"
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/errno"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.URL.Query().Get("token")

		// 无token字段
		if tokenStr == "" {
			handler.SendBaseResp(c, errno.LoginErr)
			c.Abort()
			return
		}

		// token 解析
		uid, iss, err := jwtutil.ParseToken(tokenStr)
		if err != nil {
			handler.SendBaseResp(c, errno.LoginErr)
			c.Abort()
			return
		}
		// 校验发布者
		if iss != constants.JwtIssuer {
			handler.SendBaseResp(c, errno.LoginErr)
			c.Abort()
			return
		}
		// 通过校验
		c.Set("uid", uid)
	}
}
