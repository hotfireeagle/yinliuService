package model

import "time"

/**
** 用户表定义
 */
type User struct {
	Phone    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

/**
** 用户token对象定义
 */
type UserToken struct {
	Uid string    `json:"uid"`
	Exp time.Time `json:"exp"`
}
