package model

// 按钮类型
type Button struct {
	BaseTable
	Icon        string `json:"icon" validate:"required"`
	Title       string `json:"title" validate:"required"`
	DescTxt        string `json:"descTxt" validate:"required"`
	BtnTxt      string `json:"btnTxt" validate:"required"`
	RedirectUrl string	`json:"redirectUrl" validate:"required"`
}
