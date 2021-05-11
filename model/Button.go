package model

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
	Created     int64 `gorm:"autoCreateTime"`
	Updated     int64 `gorm:"autoUpdateTime"`
	Apps        []App `gorm:"many2many:app_buttons"`
}
