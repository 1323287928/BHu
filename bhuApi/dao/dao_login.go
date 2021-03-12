package dao

import (
	"bhuApi/model"
	"fmt"
)

func Login(user model.User) model.User  {
	initDB()
	var u model.User
	sqlStr:="SELECT * FROM userinfo WHERE email=? AND password=?"
	err := db.Get(&u, sqlStr, user.Email, user.Password)
	if err!=nil{
		fmt.Println("出错了",err)
	}
	return u

}
