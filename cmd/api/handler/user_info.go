package handler

import "github.com/gin-gonic/gin"

func UserInfo(c *gin.Context) {
	userId:=c.DefaultQuery("user_id","Wangyu")
	token:=c.DefaultQuery("token","XiaoLiAiwo")
}
