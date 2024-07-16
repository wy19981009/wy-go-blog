package views

import (
	"errors"
	"log"
	"net/http"
	"wy-go-blog/common"
	"wy-go-blog/service"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 数据库查询

	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("首页获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！！"))
	}
	index.WriteData(w, hr)

}
