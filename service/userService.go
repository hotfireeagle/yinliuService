package service

import (
	"yinliuService/model"
	"yinliuService/tool"
)

// 利用手机号查用户
func FindUserByPhoneService(phone string) (*model.User, error) {
	stmt := "select id, phone, password from users where phone = ?"
	result := db.QueryRow(stmt, phone)
	var (
		id string
		tel string
		password string
	)
	if err := result.Scan(&id, &tel, &password); err != nil {
		return nil, err
	}
	item := model.User{
		Phone: tel,
		Password: password,
	}
	item.ID = id
	return &item, nil
}

// 创建新用户
func CreateUserService(obj *model.User) error {
	stmt := "insert into users (id, phone, password, created_at) values (?, ?, ?, ?)"
	_, err := db.Exec(stmt, obj.ID, obj.Phone, obj.Password, obj.CreatedAt)
	return err
}

// 通过id查找用户是否存在
func FindUserByIdService(id string) bool {
	stmt := "select id from users where id = ? and deleted_at is null"
	res := db.QueryRow(stmt, id)
	var i string
	err := res.Scan(&i)
	if err != nil {
		tool.ErrLog(err)
		return false
	}
	if i == "" {
		return false
	}
	return true
}

// 通过手机号查找用户是否已存在
func FindUserIsExistsByPhoneService(phone string) bool {
	stmt := "select id from users where phone = ? and deleted_at is null"
	row, _ := db.Exec(stmt, phone)
	num, _ := row.RowsAffected()
	if num != 0 {
		return true
	}
	return false
}