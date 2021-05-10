package model

/**
** 用户表定义
 */
type User struct {
	Phone    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
