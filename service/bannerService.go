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
	return db.Db.Table("app_banners").Where("banner_id = ?", id).Delete(&model.Any{})
}

func PatchBannerService(id string, val *model.Banner) *gorm.DB {
	//banner := model.Banner{Id: id}
	//return db.Db.Model(&banner).Select("Src", "RedirectUrl").Updates(*val)
	return db.Db.Table("banners").Where("id = ?", id).Select("Src", "RedirectUrl").Updates(*val)
}
