package model

/**
** Banner表定义
 */
type Banner struct {
	Id          string `gorm:"primaryKey;unique;not null"`
	Src         string `gorm:"not null"` // 图片链接
	RedirectUrl string // 点击图片后跳转的链接
	Apps        []App  `gorm:"many2many:app_banners"`
	Created     int64  `gorm:"autoCreateTime"`
	Updated     int64  `gorm:"autoUpdateTime"`
}

/**
** Banner响应数据定义
 */
type BannerJson struct {
	Id          string `json:"id"`
	Src         string `json:"src"`
	RedirectUrl string `json:"redirectUrl"`
}

type CreateBannerJson struct {
	Src         string
	RedirectUrl string
	AppIds      []string
}
