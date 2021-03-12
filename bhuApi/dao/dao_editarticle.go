package dao

func EditArticle(articleid int,icon string,title string,detail string) bool  {
	initDB()
	sqlStr:="update articleinfo set articleicon=?,articletitle=?,articledetail=? where articleid=?"
	exec, _ := db.Exec(sqlStr, icon, title, detail, articleid)
	affected, _ := exec.RowsAffected()
		if affected!=0{
			return  true
		}else {
			return false
		}
}
