package service

import (
	"errors"
	"wy-go-blog/dao"
	"wy-go-blog/models"
	"wy-go-blog/utils"
)

func Login(username, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "wy-go-blog")
	user := dao.GetUser(username, passwd)
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}
	uid := user.Uid
	// 生成token
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.Username = user.Username
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
