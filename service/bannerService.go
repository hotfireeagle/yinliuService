package service

import (
	"time"
	"yinliuService/model"
	"yinliuService/tool"
)

// 新建banner的SQL
func CreateBannerService(bannerObj *model.Banner) error {
	createBanner := `insert into banners (id,src,redirect_url,created_at) values (?, ?, ?, ?)`
	_, err := db.Exec(createBanner, bannerObj.ID, bannerObj.Src, bannerObj.RedirectUrl, bannerObj.CreatedAt)
	return err
}

// 查询所有banner的SQL
func SearchAllBannerService() (*[]model.Banner, error) {
	s := "select id, src, redirect_url, created_at from banners where deleted_at is null"
	result := make([]model.Banner, 0)
	rows, err := db.Query(s)
	if err != nil {
		return &result, err
	}
	defer rows.Close()
	for rows.Next() {
		item := new(model.Banner)
		if err := rows.Scan(item.ID, item.Src, item.RedirectUrl, item.CreatedAt); err != nil {
			tool.ErrLog(err)
			continue
		}
		result = append(result, *item)
	}
	return &result, nil
}

// 删除banner的处理
func DelBannerService(id string) error {
	now := time.Now()
	s := "update banners set deleted_at = ? where id = ?"
	_, err := db.Exec(s, now, id)
	return err
}

// 更新banner的处理
func UpdateBannerService(banner *model.Banner) error {
	now := time.Now()
	stmt := "update banners set src=?, redirect_url=?, updated_at=? where id=?"
	_, err := db.Exec(stmt, banner.Src, banner.RedirectUrl, now, banner.ID)
	return err
}