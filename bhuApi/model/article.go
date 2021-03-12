package model
type Article struct {
	Articleid int64  `json:"articleid"`
	Userid int64 `json:"userid"`
	Articleicon string `json:"articleicon"`
	Articletitle string `json:"articletitle"`
	Articledetail string `json:"articledetail"`
}