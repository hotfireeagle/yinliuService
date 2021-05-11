package model

/**
** app表定义
 */
type App struct {
	Id      string   `gorm:"primaryKey;unique;not null"`
	Name    string   // 应用名
	Banners []Banner `gorm:"many2many:app_banners"` // banner列表
	Menus   []Menu   `gorm:"many2many:app_menus"`   // 菜单栏列表
	Buttons []Button `gorm:"many2many:app_buttons"` // 按钮列表
}
