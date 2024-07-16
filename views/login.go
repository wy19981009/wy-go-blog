package views

import (
	"net/http"
	"wy-go-blog/common"
	"wy-go-blog/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}
