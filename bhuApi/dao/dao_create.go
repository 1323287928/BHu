package dao

import "bhuApi/model"

func Create(article model.Article) bool{
	initDB()
	sqlStr:="insert into articleinfo(userid,articleicon,articletitle,articledetail) values(?,?,?,?)"
	exec, _ := db.Exec(sqlStr, article.Userid, article.Articleicon, article.Articletitle, article.Articledetail)
	affected, _ := exec.RowsAffected()
	if affected!=0{
		return true
	}else {
		return false
	}

}
