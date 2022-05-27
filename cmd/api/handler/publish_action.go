package handler

import (
	"context"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/errno"

	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	fileData, err := file.Open()
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	rawData := make([]uint8, file.Size)
	_, err = fileData.Read(rawData)
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
	}
	data := make([]int8, file.Size)
	for i, b := range rawData {
		data[i] = int8(b)
	}
	req := video.PublishActionRequest{
		Title: c.PostForm("title"),
		Data:  data,
	}
	err = rpc.PublishAction(context.Background(), &req)
	if err != nil {
		sendBaseResp(c, errno.ConvertErr(err))
		return
	}
	sendBaseResp(c, errno.Success)
}
