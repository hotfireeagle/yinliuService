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

/**
** 多对多，根据应用ID找出所有关联的bannerID
 */
func FindRelatedBannersByAppId(appId string) []string {
	var result []string
	db.Db.Table("app_banners").Select("banner_id").Where("app_id = ?", appId).Find(&result)
	return result
}
