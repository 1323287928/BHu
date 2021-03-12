package dao

import "fmt"

func EditColumnIntro(userid int, columnname string, columnintro string) bool   {
	initDB()
	sqlStr:="update userinfo set columnname=?,columnintro=? where userid=?"
	exec, err := db.Exec(sqlStr, columnname, columnintro, userid)
	if err!=nil{
		fmt.Println("出错了",err)
	}
	affected, err := exec.RowsAffected()
	if err!=nil{
		fmt.Println("出错了",err)
	}
	if affected!=0{
		return true
	}else {
		return  false
	}

}
