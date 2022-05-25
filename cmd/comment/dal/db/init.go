package db

import (
	"douyin-micro/pkg/constants"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	DB *gorm.DB
)

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(errors.New("fail to connect to the DB"))
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(errors.New("fail to use open-tracing"))
	}
	err = DB.AutoMigrate(&Comment{})
	if err != nil {
		panic(errors.New("fail to migrate the Comment table"))
	}
}
