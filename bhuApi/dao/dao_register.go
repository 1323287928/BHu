package dao

import (
	"bhuApi/model"
	"fmt"
)


func queryIsExist(user model.User) bool  {
  var u model.User
	sqlStr:="select * from userinfo where email=?"
	_= db.Get(&u, sqlStr, string(user.Email))

	if u.Username!=""{
		fmt.Println(u,"存在")
	  return true
	}else {
		fmt.Println("不存在")
		return false
	}


}
func Register(user model.User) (bool,bool) {
	initDB()
	sqlStr:="insert into userinfo(email,username,password,age) values(?,?,?,?)"
	isExist := queryIsExist(user)

	if isExist {
		return isExist,false
	}else {
		exec, _ := db.Exec(sqlStr, user.Email, user.Username, user.Password, user.Age)
		affected, _ := exec.RowsAffected()
		if affected!=0{
			fmt.Println(affected)
			return isExist,true

		}else {
			fmt.Println(affected)
			return isExist,false
		}


	}


}
