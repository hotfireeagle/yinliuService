package model

import "time"

// 用户类型
type User struct {
	BaseTable
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// 用户token定义
type UserToken struct {
	Uid string    `json:"uid"`
	Exp time.Time `json:"exp"`
}
