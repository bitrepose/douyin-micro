package handler

import (
	"douyin-micro/kitex_gen/video"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	userId:=c.DefaultQuery("user_id","Wangyu")
	token:=c.DefaultQuery("token","nothing...")
	videoId:=c.DefaultQuery("video_id","-1")
	actionType:=c.DefaultQuery("action_type","0")
}
