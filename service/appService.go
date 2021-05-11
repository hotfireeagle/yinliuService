package service

import (
	"gorm.io/gorm"
	"yinliuService/db"
	"yinliuService/model"
)

/**
** 创建App
 */
func CreateAppService(app *model.App) *gorm.DB {
	return db.Db.Create(app)
}

/**
** 根据appid查找出所有应用列表
 */
func FindAppsByIds(ids []string) *[]model.App {
	var result []model.App
	db.Db.Where("id IN ?", ids).Find(&result)
	return &result
}
