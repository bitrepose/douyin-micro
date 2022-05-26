package main

import (
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/tracer"

	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)

}

func main() {
	Init()
	r := gin.New()
	// authMiddleware, _ = jwt

	v1 := r.Group("/v1/douyin")

	user1 := v1.Group("/user")
	// 用户注册
	user1.POST("/register")
	// 用户登陆
	user1.POST("/login")
	// 用户信息
	user1.GET("/")

	// 视频流
	v1.GET("/feed")

	publish1 := v1.Group("/publish")
	// 投稿
	publish1.POST("/action")
	// 发布列表
	publish1.GET("/list")

	fav1 := v1.Group("/favorite")
	// 点赞
	fav1.POST("/action")
	// 点赞列表
	fav1.GET("/list")

	comment1 := v1.Group("/comment")
	// 评论
	comment1.POST("/action")
	// 评论列表
	comment1.GET("/list")

	relation1 := v1.Group("/relation")
	// 关注
	relation1.POST("/action")
	// 关注列表
	relation1.GET("/follow/list")
	// 粉丝列表
	relation1.GET("/follower/list")
}
