package common

import (
	"sync"
	"wy-go-blog/config"
	"wy-go-blog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		// 耗时操作
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		// 标识已完成
		w.Done()
	}()
}
