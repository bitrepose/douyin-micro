package handler

import (
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	var req RelationReq
	if err := c.ShouldBind(&req); err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
		return
	}
	if req.ActionType < 1 || req.ActionType > 2 || req.UserId <= 0 || req.ToUserId <= 0 {
		sendBaseResp(c, errno.ParamErr)
		return
	}

	//token 操作
	var userId int64

	var request = user.RelationActionRequest{userId, req.ToUserId, req.ActionType}
	err := rpc.RelationAction(c, request)
	sendBaseResp(c, err)

}

type RelationReq struct {
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
	ToUserId   int64  `json:"to_user_id"`
	ActionType int32  `json:"action_type"`
}
