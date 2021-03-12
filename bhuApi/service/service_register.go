package service

import (
	"bhuApi/dao"
	"bhuApi/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)




var u model.User
var article model.Article
var isLogin bool
var isEdit bool
var isDel bool
func Register()  {

	r:=gin.Default()


	r.GET("register", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		email:=c.Query("email")
		username:=c.Query("username")
		password:=c.Query("password")
		age:=c.Query("age")

		u.Email=email
		u.Username=username
		u.Password=password
		atoi, _ := strconv.ParseInt(age,10,64)
		u.Age=atoi

		isExist,isRegister := dao.Register(u)
		fmt.Println(isExist)
		c.JSON(http.StatusOK,gin.H{
			"email":email,
			"username":username,
			"password":password,
			"age":age,
			"exits":isExist,
			"isRegister":isRegister,

		})


	})

	r.GET("login", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		email:=c.Query("email")

		password:=c.Query("password")


		u.Email=email

		u.Password=password

		loginUser := dao.Login(u)
		if loginUser.Email!=""{
			isLogin=true
		}else {
			isLogin=false
		}

		c.JSON(http.StatusOK,gin.H{
			"user":loginUser,
			"isLogin":isLogin,
		})


	})
	
	r.GET("/create", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		userid := c.Query("userid")
		articleicon := c.Query("articleicon")
		articletitle := c.Query("articletitle")
		articledetail := c.Query("articledetail")
		article.Userid,_= strconv.ParseInt(userid,10,64)
		article.Articleicon=articleicon
		article.Articletitle=articletitle
		article.Articledetail=articledetail
		isCreate := dao.Create(article)
		c.JSON(http.StatusOK,gin.H{
			"created":isCreate,
		})
	})
	r.GET("/column", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")


		userid,_:=strconv.Atoi(c.Query("userid"))
		columns := dao.Column(userid)
		c.JSON(http.StatusOK,gin.H{
			"columns":columns,
		})


	})

	r.GET("/userinfo", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		userid,_:=strconv.Atoi(c.Query("userid"))

		 u=dao.UserInfo(userid)
		 c.JSON(http.StatusOK,u)
	})

	r.GET("/editarticle", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		articleid,_:=strconv.Atoi(c.Query("articleid"))
		articledetail:=c.Query("articledetail")
		articleicon:=c.Query("articleicon")
			articletitle:=c.Query("articletitle")
		isEdit=dao.EditArticle(articleid,articleicon,articletitle,articledetail)
		c.JSON(http.StatusOK,gin.H{
			"isEdit":isEdit,
		})
	})

	r.GET("/delarticle", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		articleid,_:=strconv.Atoi(c.Query("articleid"))

		isDel=dao.DelArticle(articleid)
		c.JSON(http.StatusOK,gin.H{
			"isDel":isDel,
		})
	})

	r.GET("/editpersonalintro", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		userid,_:=strconv.Atoi(c.Query("userid"))
		username:=c.Query("username")
		personalintro:=c.Query("personalintro")
		icon:=c.Query("icon")
		 isEditPersonal := dao.EditPersonalIntro(userid,icon, username, personalintro)
		 c.JSON(http.StatusOK,gin.H{
		 	"isEditPersonalIntro":isEditPersonal,
		 })

	})

	r.GET("/editcolumnintro", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		userid,_:=strconv.Atoi(c.Query("userid"))
		columnname:=c.Query("columnname")
		columnintro:=c.Query("columnintro")
		isEditColumn := dao.EditColumnIntro(userid, columnname, columnintro)
		c.JSON(http.StatusOK,gin.H{
			"isEditColumnIntro":isEditColumn,
		})

	})
	r.GET("/loadmove", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		//允许请求头修改的类容
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//允许使用cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		var move []model.LoadMoveUser
		move= dao.LoadMove()
		c.JSON(http.StatusOK,gin.H{
			"moveuser":move,
		})
	})


	r.Run()
}

