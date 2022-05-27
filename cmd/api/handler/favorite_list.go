package handler

import "github.com/gin-gonic/gin"

func FavoriteList(c *gin.Context) {
	userId:=c.DefaultQuery("user_id","Wangyu")
	toekn:=c.DefaultQuery("token","nothing...")
}
