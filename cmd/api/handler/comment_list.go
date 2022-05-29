package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/comment"
	"douyin-micro/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommentList(c *gin.Context) {
	var params struct {
		Token   string `json:"token" form:"token"`
		VideoId int64  `json:"video_id" form:"video_id"`
	}
	if err := c.BindQuery(&params); err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	if params.VideoId < 0 {
		sendBaseResp(c, errno.ParamErr)
	}
	req := comment.CommentListRequest{
		VideoId: params.VideoId,
	}
	resp, err := rpc.CommentList(context.Background(), &req)
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}

	c.JSON(http.StatusOK, resp)
}
