package model

import "time"

/**
*** 中间菜单区数据库的定义
 */
type Menu struct {
	Id          string `gorm:"primaryKey;unique;not null"`
	Icon        string `gorm:"not null"`
	Text        string `gorm:"not null"`
	RedirectUrl string
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time `gorm:"autoUpdateTime"`
	Apps        []App     `gorm:"many2many:app_menus"`
}

/**
** 新增menu时的json参数定义
 */
type CreateNewMenuJson struct {
	Icon        string
	Text        string
	RedirectUrl string
	AppIds      []string
}

/**
** 单个menu的json对象描述
 */
type MenuJson struct {
	Id          string    `json:"id"`
	Icon        string    `json:"icon"`
	Text        string    `json:"text"`
	RedirectUrl string    `json:"redirectUrl"`
	Created     time.Time `json:"created"`
}
