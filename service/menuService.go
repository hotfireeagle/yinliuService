package service

import (
	"gorm.io/gorm"
	"yinliuService/db"
	"yinliuService/model"
)

/**
** 创建新banner
 */
func CreateNewMenuService(menu *model.Menu) *gorm.DB {
	return db.Db.Create(menu)
}

/**
** 根据menuId列表找出所有menu列表
 */
func FindMenusByIds(ids []string) *[]model.Menu {
	var result []model.Menu
	db.Db.Where("id IN ?", ids).Order("updated desc").Find(&result)
	return &result
}

/**
** 删除menu
 */
func DeleteMenuService(id string) *gorm.DB {
	return db.Db.Table("app_menus").Where("menu_id = ?", id).Delete(&model.Any{})
}

/**
** 更新menu
 */
func PatchMenuService(id string, val *model.Menu) *gorm.DB {
	//banner := model.Banner{Id: id}
	//return db.Db.Model(&banner).Updates(*val)
	return db.Db.Table("menus").Where("id = ?", id).Updates(*val)
}
