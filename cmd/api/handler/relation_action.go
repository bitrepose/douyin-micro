package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	var req RelationReq
	if err := c.ShouldBind(&req); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
		return
	}
	if req.ActionType < 1 || req.ActionType > 2 || req.ToUserId <= 0 {
		SendBaseResp(c, errno.ParamErr)
		return
	}

	//token 操作
	var userId = c.GetInt64("uid")

	var request = user.RelationActionRequest{UserId: userId, ToUserId: req.ToUserId, ActionType: req.ActionType}
	err := rpc.RelationAction(context.Background(), request)
	SendBaseResp(c, err)
}

type RelationReq struct {
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}
