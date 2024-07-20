package server

import (
	"log"
	"net/http"
	"wy-go-blog/config"
	"wy-go-blog/router"
)

var App = &MsServer{}

type MsServer struct {
}

func (*MsServer) Start() {
	ip := config.Cfg.System.ServerURL
	port := config.Cfg.System.ServerPort
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
