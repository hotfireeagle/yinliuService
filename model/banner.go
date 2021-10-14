package model

// Banner类型
type Banner struct {
	BaseTable
	Src         string    `json:"src" validate:"required"`
	RedirectUrl string    `json:"redirectUrl" validate:"required"`
}
