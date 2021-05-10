package service

import (
	"gorm.io/gorm"
	"yinliuService/db"
	"yinliuService/model"
)

/**
*** 利用手机号查用户
 */
func FindUserByPhoneService(email string) *model.User {
	result := model.User{}
	db.Db.Where("phone = ?", email).First(&result)
	return &result
}

/**
** 判断该邮箱在数据库中是否已存在
 */
func IsPhoneExists(property string, val string) bool {
	propertyStr := property + "= ?"
	result := model.User{}
	db.Db.Where(propertyStr, val).First(&result)
	return result.Password != ""
}

/**
** 创建用户
 */
func CreateUserService(obj *model.User) *gorm.DB {
	return db.Db.Create(obj)
}
