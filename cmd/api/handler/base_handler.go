package handler

import (
	"douyin-micro/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func SendBaseResp(c *gin.Context, err errno.ErrNo) {
	c.JSON(http.StatusOK, Response{
		StatusCode: int(err.ErrCode),
		StatusMsg:  err.ErrMsg,
	})
}
