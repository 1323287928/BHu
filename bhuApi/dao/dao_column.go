package dao

import "bhuApi/model"

func Column(userid int) []model.Article {
	initDB()
	sqlStr:="select * from articleinfo where userid=?"
	var article []model.Article
	db.Select(&article,sqlStr,userid)
	return  article

}
