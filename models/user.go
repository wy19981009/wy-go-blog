package models

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"userName"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	Username string `json:"userName"`
	Avatar   string `json:"avatar"`
}
