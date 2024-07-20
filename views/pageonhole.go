package views

import (
	"net/http"
	"wy-go-blog/common"
	"wy-go-blog/service"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
