package dao

import (
	"fmt"
	"log"
	"wy-go-blog/models"
)

func GetUserNameById(id int) string {
	row := DB.QueryRow("select user_name from blog_user where uid = ?", id)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(username, passwd string) *models.User {
	fmt.Println(username, passwd)
	row := DB.QueryRow("select * from blog_user where user_name = ? and passwd=?", username, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.Username, &user.Password, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}
