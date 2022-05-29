package handler

import (
	"context"
	jwtutil "douyin-micro/cmd/api/jwt_util"
	"douyin-micro/cmd/api/rpc"
	"douyin-micro/kitex_gen/video"
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/errno"

	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	tokenStr := c.PostForm("token")
	uid, iss, err := jwtutil.ParseToken(tokenStr)
	if iss != constants.JwtIssuer || err != nil {
		SendBaseResp(c, errno.LoginErr)
	}
	file, err := c.FormFile("data")
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	fileData, err := file.Open()
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	rawData := make([]uint8, file.Size)
	_, err = fileData.Read(rawData)
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
	}
	data := make([]int8, file.Size)
	for i, b := range rawData {
		data[i] = int8(b)
	}
	req := video.PublishActionRequest{
		Title:  c.PostForm("title"),
		Data:   data,
		UserId: uid,
	}
	err = rpc.PublishAction(context.Background(), &req)
	if err != nil {
		SendBaseResp(c, errno.ConvertErr(err))
		return
	}
	SendBaseResp(c, errno.Success)
}
