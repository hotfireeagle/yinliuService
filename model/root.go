package model

import (
	"github.com/google/uuid"
	"time"
)

type ResponseCode int

const (
	Err ResponseCode = iota
	Success // 响应正常
	UnLogin                              // 未登录
	LoginOverdue                         // 登录过期
)

type IResponse struct {
	Status ResponseCode `json:"status"` // 响应状态
	Msg  string       `json:"msg"`  // 报错信息
	ErrLog string `json:"errLog"`
	Data interface{}  `json:"data"` // 正式数据
}

type BaseTable struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func (obj *BaseTable) GenerateUUID() {
	obj.ID = uuid.NewString()
	obj.CreatedAt = time.Now()
}