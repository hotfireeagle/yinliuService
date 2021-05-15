package service

import (
	"gorm.io/gorm"
	"yinliuService/db"
	"yinliuService/model"
)

/**
** 创建新button
 */
func CreateNewButtonService(button *model.Button) *gorm.DB {
	return db.Db.Create(button)
}

/**
** 根据menuId列表找出所有button列表
 */
func FindButtonsByIds(ids []string) *[]model.Button {
	var result []model.Button
	db.Db.Where("id IN ?", ids).Order("updated desc").Find(&result)
	return &result
}

/**
** 删除button
 */
func DeleteButtonService(id string) *gorm.DB {
	return db.Db.Table("app_buttons").Where("button_id = ?", id).Delete(&model.Any{})
}

/**
** 更新button
 */
func PatchButtonService(id string, val *model.Button) *gorm.DB {
	button := model.Button{Id: id}
	return db.Db.Model(&button).Updates(*val)
}
