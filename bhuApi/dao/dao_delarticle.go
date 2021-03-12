package dao

func DelArticle(articleid int)  bool {
		initDB()
		sqlStr:="DELETE  FROM articleinfo WHERE articleid=?"
	exec, _ := db.Exec(sqlStr, articleid)
	affected, _ := exec.RowsAffected()
	if affected!=0{
		return true
	}else {
		return  false
	}

}
