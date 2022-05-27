package handler

import "github.com/gin-gonic/gin"

func CommentList(c *gin.Context) {
	token:=c.DefaultQuery("token","nothing...")
	videoId:=c.DefaultQuery("video_id","-1")
}
