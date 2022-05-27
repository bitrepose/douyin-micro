package handler

import "github.com/gin-gonic/gin"

func FollowList(c *gin.Context) {
	userId:=c.DefaultQuery("user_id","Wangyu")
	token:=c.DefaultQuery("token","nothing...")
	
}
