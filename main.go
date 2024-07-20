package main

import (
	"wy-go-blog/common"
	"wy-go-blog/server"
)

func init() {
	// 模版加载
	common.LoadTemplate()
}

func main() {
	// 程序入口，一个项目 只能有一个 main.go 入口
	// web程序
	server.App.Start()
}
