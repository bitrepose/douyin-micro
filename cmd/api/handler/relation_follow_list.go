package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowList(c *gin.Context) {
	var req FollowRequest
	if err := c.ShouldBind(&req); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
		return
	}
	if req.UserId <= 0 {
		SendBaseResp(c, errno.ParamErr)
		return
	}
	//token操作
	var userId = c.GetInt64("uid")

	var request = user.RelationFollowListRequest{UserId: userId, ReqUserId: req.UserId}
	users, err := rpc.RelationFollowList(context.Background(), request)
	c.JSON(http.StatusOK, FollowResponse{Response{StatusMsg: err.ErrMsg, StatusCode: int(err.ErrCode)}, users})
}

type FollowRequest struct {
	UserId int64 `json:"user_id" form:"user_id"`
}

type FollowResponse struct {
	Response
	UserList []*user.User `json:"user_list"`
}
