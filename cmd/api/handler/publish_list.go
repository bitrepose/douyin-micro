package handler

import "github.com/gin-gonic/gin"

func PublishList(c *gin.Context) {
	token:=c.DefaultQuery("token","nothing...")
	userId:=c.DefaultQuery("user_id","Wangyu")
}
