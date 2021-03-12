package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var db *sqlx.DB
func initDB() (err error) {
	dsn:="root:123456@tcp(127.0.0.1:3306)/bhu?charset=utf8mb4&parseTime=True"
	db,err=sqlx.Connect("mysql",dsn)
	if err!=nil{
		fmt.Printf("connect DB failed,err:%v\n ",err)
		return
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return err


}

