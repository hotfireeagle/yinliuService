package service

import (
	"gorm.io/gorm"
	"yinliuService/db"
	"yinliuService/model"
)

/**
** 创建App
 */
func CreateBannerService(banner *model.Banner) *gorm.DB {
	return db.Db.Create(banner)
}

func FindBannersByIds(ids []string) *[]model.Banner {
	var result []model.Banner
	db.Db.Where("id IN ?", ids).Find(&result)
	return &result
}
