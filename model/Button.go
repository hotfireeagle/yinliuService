package model

import "time"

/**
** 底部按钮库的定义
 */
type Button struct {
	Id          string `gorm:"primaryKey;unique;not null"`
	Icon        string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Desc        string `gorm:"not null"`
	BtnTxt      string `gorm:"not null"`
	RedirectUrl string
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time `gorm:"autoUpdateTime"`
	Apps        []App     `gorm:"many2many:app_buttons"`
}

/**
** 新增menu时的json参数定义
 */
type CreateNewButtonJson struct {
	Icon        string
	Title       string
	Desc        string
	BtnTxt      string
	RedirectUrl string
	AppIds      []string
}

/**
** 单个menu的json对象描述
 */
type ButtonJson struct {
	Id          string    `json:"id"`
	Icon        string    `json:"icon"`
	Title       string    `json:"title"`
	Desc        string    `json:"desc"`
	BtnTxt      string    `json:"btnTxt"`
	RedirectUrl string    `json:"redirectUrl"`
	Created     time.Time `json:"created"`
}
