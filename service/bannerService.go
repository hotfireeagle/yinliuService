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
