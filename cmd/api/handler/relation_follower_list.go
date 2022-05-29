package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/user"
	"douyin-micro/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowerList(c *gin.Context) {
	var req FollowerRequest
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

	users, err := rpc.RelationFollowerList(context.Background(), user.RelationFollowerListRequest{UserId: userId, ReqUserId: req.UserId})
	c.JSON(http.StatusOK, FollowerResponse{Response{StatusCode: int(err.ErrCode), StatusMsg: err.ErrMsg}, users})
}

type FollowerRequest struct {
	UserId int64 `json:"user_id"`
}

type FollowerResponse struct {
	Response
	UserList []*user.User `json:"user_list"`
}
