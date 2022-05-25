package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:SUDAcs647#SQL@tcp(139.196.145.51:3306)/douyin?parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err.(any))
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err.(any))
	}

	if err = DB.AutoMigrate(&User{}, &FollowRelation{}); err != nil {
		panic(err.(any))
	}

	//m := DB.Migrator()
	//if m.HasTable(&User{}) {
	//} else {
	//	if err = m.CreateTable(&User{}); err != nil {
	//		panic(err.(any))
	//	}
	//}
	//
	//if m.HasTable(&FollowRelation{}) {
	//} else {
	//	if err = m.CreateTable(&FollowRelation{}); err != nil {
	//		panic(err.(any))
	//	}
	//}

}
