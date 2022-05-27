package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	last_time:=c.DefaultQuery("last_time",time.Now().String())
	
}
