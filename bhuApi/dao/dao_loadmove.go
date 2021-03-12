package dao

import (
	"bhuApi/model"
	"fmt"
)

func LoadMove()   []model.LoadMoveUser {
	initDB()
	sqlStr:="select userid,username,columnintro from userinfo where userid>?"
	var users []model.LoadMoveUser
	err:=db.Select(&users,sqlStr,0)
	fmt.Println(users)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}

	return  users
}
