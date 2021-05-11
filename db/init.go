package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"yinliuService/model"
)

var Db *gorm.DB

func init() {
	dsn := "yinliu:a6HLLshzrH7KbLk2@tcp(120.76.196.216:3306)/yinliu?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接错误")
	}

	err = Db.AutoMigrate(&model.User{})
	err = Db.AutoMigrate(&model.App{})
	err = Db.AutoMigrate(&model.Banner{})
	err = Db.AutoMigrate(&model.Menu{})
	err = Db.AutoMigrate(&model.Button{})

	if err != nil {
		panic("创建数据表错误")
	}
}
