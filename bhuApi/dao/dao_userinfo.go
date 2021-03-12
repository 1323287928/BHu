package dao

import (
	"bhuApi/model"
	"fmt"
)

func UserInfo(userid int)  model.User {
	initDB()
	sqlStr:="select * from userinfo where userid=?"
	var u model.User
	err := db.Get(&u, sqlStr, userid)
	if err!=nil{
		fmt.Println("查询出错了")

	}
	return  u

}
