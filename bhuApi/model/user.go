package model

type User struct {
	Userid int64 `json:"userid" sql:"userid"`
	Email  string `json:"email" sql:"email"`
	Username string `json:"username" sql:"username"`
	Password string `json:"password" sql:"password"`
	Age int64 `json:"age" sql:"password"`
	Icon string `json:"icon" sql:"password"`
	Columnname string `json:"columnname" sql:"columnname"`
	Columnintro string `json:"columnintro" sql:"columnintro"`
	Personalintro string `json:"personalintro" sql:"personalintro" `
}
