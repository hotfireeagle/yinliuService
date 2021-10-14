package service

import (
	"time"
	"yinliuService/model"
	"yinliuService/tool"
)

// 新建menu
func CreateNewMenuService(menu *model.Menu) error {
	stmt := "insert into menus (id, icon, text, redirect_url, created_at) values (?, ?, ?, ?, ?)"
	_, err := db.Exec(stmt, menu.ID, menu.Icon, menu.Text, menu.RedirectUrl, menu.CreatedAt)
	return err
}

// 查找出所有的menu
func FindMenusService() (*[]model.Menu, error) {
	result := make([]model.Menu, 0)
	stmt := "select id, icon, text, redirect_url, created_at from menus where deleted_at is null order by created_at desc"
	rows, err := db.Query(stmt)
	if err != nil {
		return &result, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id string
			icon string
			text string
			redirectUrl string
			createdAt time.Time
		)
		err := rows.Scan(&id, &icon, &text, &redirectUrl, &createdAt)
		if err != nil {
			tool.ErrLog(err)
			continue
		}
		item := model.Menu{
			Icon: icon,
			Text: text,
			RedirectUrl: redirectUrl,
		}
		item.ID = id
		item.CreatedAt = createdAt
		result = append(result, item)
	}
	return &result, nil
}

// 删除menu
func DeleteMenuService(id string) error {
	now := time.Now()
	stmt := "update menus set deleted_at = ? where id = ?"
	_, err := db.Exec(stmt, now, id)
	return err
}

// 更新menu
func PatchMenuService(menu *model.Menu) error {
	now := time.Now()
	stmt := "update menus set icon=?, text=?, redirect_url=?, updated_at=? where id=?"
	_, err := db.Exec(stmt, menu.Icon, menu.Text, menu.RedirectUrl, now, menu.ID)
	return err
}
