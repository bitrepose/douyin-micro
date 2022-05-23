package db

import (
	"douyin-micro/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			// NamingStrategy: schema.NamingStrategy{
			// 	SingularTable: true, // Migrate时创建单数Video表
			// },
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(new(Video), new(FavoriteRelation)); err != nil {
		panic(err)
	}
	// m := DB.Migrator()
	// if m.HasTable(&Video{}) {
	// 	return
	// }
	// if err = m.CreateTable(&Video{}); err != nil {
	// 	panic(err)
	// }
}
