package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

// GetRequestJsonParam 获取JSON参数
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
