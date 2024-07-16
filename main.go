package main

import (
	"log"
	"net/http"
	"wy-go-blog/common"
	"wy-go-blog/router"
)

func init() {
	// 模版加载
	common.LoadTemplate()
}

func main() {
	// 程序入口，一个项目 只能有一个 main.go 入口
	// web程序
	service := http.Server{
		Addr: "127.0.0.1:7888",
	}

	//路由
	router.Router()
	if err := service.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
