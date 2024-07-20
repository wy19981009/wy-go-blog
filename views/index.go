package views

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！！"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	// 每页显示的数量
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！！"))
	}
	index.WriteData(w, hr)

}
