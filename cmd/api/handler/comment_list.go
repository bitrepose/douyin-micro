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
		VideoId int64 `json:"video_id" form:"video_id"`
	}
	if err := c.BindQuery(&params); err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	if params.VideoId < 0 {
		SendBaseResp(c, errno.ParamErr)
	}
	req := comment.CommentListRequest{
		VideoId:   params.VideoId,
		ReqUserId: c.GetInt64("uid"),
	}
	resp, err := rpc.CommentList(context.Background(), &req)
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}

	c.JSON(http.StatusOK, resp)
}
