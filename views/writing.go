package views

import (
	"net/http"
	"wy-go-blog/common"
	"wy-go-blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
