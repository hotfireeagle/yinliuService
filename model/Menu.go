package model

/**
*** 中间菜单区数据库的定义
 */
type Menu struct {
	Id          string `gorm:"primaryKey;unique;not null"`
	Icon        string `gorm:"not null"`
	Text        string `gorm:"not null"`
	RedirectUrl string
	Created     int64 `gorm:"autoCreateTime"`
	Updated     int64 `gorm:"autoUpdateTime"`
	Apps        []App `gorm:"many2many:app_menus"`
}
