package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	lastest_time := c.Query("lastest_time")
	token := c.Query("token")
	req := video.FeedRequset{}
	if len(lastest_time) != 0 {
		lt, err := strconv.Atoi(lastest_time)
		if err != nil {
			sendBaseResp(c, errno.ConvertErr(err))
		}
		t := int64(lt)
		req.LatestTime = &t
	}

	if len(token) != 0 {

	}
	resp, err := rpc.Feed(context.Background(), &req)
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	c.JSON(http.StatusOK, resp)
}
