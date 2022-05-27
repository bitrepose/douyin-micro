package handler

import "github.com/gin-gonic/gin"

func RelationAction(c *gin.Context) {
	userId := c.DefaultQuery("user_id", "Wangyu")
	token := c.DefaultQuery("token", "nothing...")
	toUserId := c.DefaultQuery("to_user_id", "nothing...")
	actionType := c.DefaultQuery("action_type", "0")
}
