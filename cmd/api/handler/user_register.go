package handler

import "github.com/gin-gonic/gin"

func UserRegister(c *gin.Context) {
	username := c.DefaultQuery("username", "Wangyu")
	password := c.DefaultQuery("password", "WoAiXiaoLi")

}
