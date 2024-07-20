package router

import (
	"net/http"
	"wy-go-blog/api"
	"wy-go-blog/views"
)

func Router() {
	// 1.页面 views 2.数据 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
