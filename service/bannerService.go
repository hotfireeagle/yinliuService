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
	db.Db.Where("id IN ?", ids).Order("created desc").Find(&result)
	return &result
}

func DeleteBannerService(id string) *gorm.DB {
	return db.Db.Where("id = ?", id).Delete(&model.Banner{})
}

func PatchBannerService(id string, val *model.Banner) *gorm.DB {
	banner := model.Banner{Id: id}
	return db.Db.Model(&banner).Updates(*val)
}
