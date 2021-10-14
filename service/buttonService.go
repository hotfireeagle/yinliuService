package service

import (
	"time"
	"yinliuService/model"
	"yinliuService/tool"
)

// 创建新button
func CreateNewButtonService(button *model.Button) error {
	stmt := "insert into buttons(id, icon, title, desc_txt, btn_txt, redirect_url, created_at) values(?, ?, ?, ?, ?, ?, ?)"
	button.CreatedAt = time.Now()
	_, err := db.Exec(stmt, button.ID, button.Icon, button.Title, button.DescTxt, button.BtnTxt, button.RedirectUrl, button.CreatedAt)
	return err
}

// 查找所有的按钮
func FindButtons() (*[]model.Button, error) {
	var result []model.Button
	stmt := "select id, icon, title, desc_txt, btn_txt, redirect_url, created_at from buttons where deleted_at is null"
	rows, err := db.Query(stmt)
	if err != nil {
		return &result, err
	}
	defer rows.Close()
	for rows.Next() {
		item := new(model.Button)
		if err := rows.Scan(item.ID, item.Icon, item.DescTxt, item.BtnTxt, item.RedirectUrl, item.CreatedAt); err != nil {
			tool.ErrLog(err)
			continue
		}
		result = append(result, *item)
	}
	return &result, nil
}

// 删除button
func DeleteButtonService(id string) error {
	now := time.Now()
	stmt := "update buttons set deleted_at = ? where id = ?"
	_, err := db.Exec(stmt, now, id)
	return err
}

// 更新button
func PatchButtonService(button *model.Button) error {
	now := time.Now()
	stmt := "update buttons set icon=?, title=?, desc_txt=?, btn_txt=?, redirect_url=?, updated_at=? where id=?"
	_, err := db.Exec(stmt, button.Icon, button.Title, button.DescTxt, button.BtnTxt, button.RedirectUrl, now, button.ID)
	return err
	return nil
}
