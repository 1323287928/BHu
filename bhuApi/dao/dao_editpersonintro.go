package dao

import "fmt"

func EditPersonalIntro(userid int,icon string,username string,personalintro string) bool  {
	initDB()
	sqlStr:="update userinfo set icon=?,username=?,personalintro=? where userid=?"
	exec, err := db.Exec(sqlStr, icon, username, personalintro,userid)
	if err!=nil{
		fmt.Println("出错了===>",err)
	}
	affected, err := exec.RowsAffected()
	if err!=nil{
		fmt.Println("出错了===>",err)
	}
		if affected!=0{
			return  true
		}else{
			return  false
		}
}
