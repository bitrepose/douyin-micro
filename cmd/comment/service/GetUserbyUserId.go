package service
import (
	"douyin-micro/kitex_gen/user"
)
func GetUserByUserId(userId int ) *user.User{
	return &user.User{
		Id:1,
		Name:"goodman",
	}
}
