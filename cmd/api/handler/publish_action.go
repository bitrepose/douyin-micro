package handler

import (
	"douyin-micro/pkg/errno"

	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	// req := video.
}
