package handler

import "github.com/gin-gonic/gin"

func CommentAction(c *gin.Context) {
	userId:=c.DefaultQuery("user_id","Wangyu")
	token:=c.DefaultQuery("token","nothing...")
	videoId:=c.DefaultQuery("video_id","-1")
	actionType:=c.DefaultQuery("action_type","0")
	commentText:=c.DefaultQuery("comment_text","TianMiMi")
	commentId:=c.DefaultQuery("comment_id","-1")
}
