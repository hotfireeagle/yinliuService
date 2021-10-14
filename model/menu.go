package model

type Menu struct {
	BaseTable
	Icon        string `json:"icon" validate:"required"`
	Text        string `json:"text" validate:"required"`
	RedirectUrl string `json:"redirectUrl" validate:"required"`
}