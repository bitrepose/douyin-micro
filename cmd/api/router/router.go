package router

import (
	"douyin-micro/cmd/api/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() error {
	r := gin.Default()
	douyinGroup := r.Group("/douyin")
	{
		douyinGroup.GET("/feed", handler.Feed)
		userGroup := douyinGroup.Group("/user")
		{
			userGroup.POST(("/register"), handler.UserRegister)
			userGroup.POST("/login", handler.UserLogin)
			userGroup.GET("/", handler.UserInfo)
		}
		publishGroup := douyinGroup.Group("/publish")
		{
			publishGroup.POST("/action", handler.PublishAction)
			publishGroup.GET("/list", handler.PublishList)
		}
		favoriteGroup := douyinGroup.Group("/favorite")
		{
			favoriteGroup.POST("/action", handler.FavoriteAction)
			favoriteGroup.GET("/list", handler.FavoriteList)
		}
		commentGroup := douyinGroup.Group("/comment")
		{
			commentGroup.POST("/action", handler.CommentAction)
			commentGroup.GET("/list", handler.CommentList)
		}
		relationGroup := douyinGroup.Group("/relation")
		{
			relationGroup.POST("/action", handler.RelationAction)
			relationGroup.GET("/follow/list", handler.FollowList)
			relationGroup.GET("/follower/list", handler.FollowerList)
		}
	}
	err := r.Run(":9090")
	return err
}
