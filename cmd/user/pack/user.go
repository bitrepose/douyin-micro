package pack

import (
	"douyin-micro/cmd/user/dal/db"
	"douyin-micro/kitex_gen/user"
)

//ConvUser 将模型的user 转换为 响应里的user
func ConvUser(prev db.User) *user.User {
	result := &user.User{
		Id:            prev.ID,
		Name:          prev.Name,
		FollowCount:   &prev.FollowCount,
		FollowerCount: &prev.FollowerCount,
		IsFollow:      prev.IsFollow}
	return result
}
