package model

type LoadMoveUser struct {
	Userid int64 `json:"userid" sql:"userid"`
	Username string `json:"username" sql:"username"`
	Columnintro string `json:"columnintro" sql:"columnintro"`

}