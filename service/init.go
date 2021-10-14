package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"yinliuService/environmentVariable"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", *environmentVariable.DSN)

	// key point occurs error should panic
	if err != nil {
		panic(err)
	}
}