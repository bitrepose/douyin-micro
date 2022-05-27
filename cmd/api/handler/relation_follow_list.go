package handler

import (
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowList(c *gin.Context) {
	var req FollowRequest
	if err := c.ShouldBind(&req); err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
		return
	}
	if req.UserId <= 0 {
		sendBaseResp(c, errno.ParamErr)
		return
	}
	//token操作
	var userId int64

	var request = user.RelationFollowListRequest{userId, req.UserId}
	users, err := rpc.RelationFollowList(c, request)
	c.JSON(http.StatusOK, FollowResponse{Response{StatusMsg: err.ErrMsg, StatusCode: int(err.ErrCode)}, users})
}

type FollowRequest struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type FollowResponse struct {
	Response
	UserList []*user.User `json:"user_list"`
}
