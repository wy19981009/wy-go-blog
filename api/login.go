package api

import (
	"net/http"
	"wy-go-blog/common"
	"wy-go-blog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接收用户名和密码， 返回对应的json数据
	params := common.GetRequestJsonParam(r)
	username := params["username"].(string)
	passwd := params["passwd"].(string)
	LoginRes, err := service.Login(username, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, LoginRes)
}
