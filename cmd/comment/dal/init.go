package dal

import (
	"douyin-micro/cmd/comment/dal/db"
)

func Init() {
	db.InitDB()
}
