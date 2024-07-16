package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wy-go-blog/common"
	"wy-go-blog/dao"
	"wy-go-blog/models"
	"wy-go-blog/service"
	"wy-go-blog/utils"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostById(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	// 获取用户id
	token := r.Header.Get("Authorization")
	_, c, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := c.Uid
	// POST代表save操作
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		ptype := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			ptype,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		// update

	}

}
