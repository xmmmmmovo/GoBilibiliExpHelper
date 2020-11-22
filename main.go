package main

import (
	"GoBilibiliExpHelper/config"
	"GoBilibiliExpHelper/errors"
	"GoBilibiliExpHelper/http"
	"GoBilibiliExpHelper/service"
	"log"
)

func main() {
	err := config.LoadConfig("./config.yml")
	if err != nil {
		config.SaveConfig("./config.yml")
	}
	err = config.EnvReader()
	if err != nil {
		log.Println(err.Error())
		return
	}
	http.OnResponse(func(resp *map[string]interface{}) (*map[string]interface{}, error) {
		data := * resp
		log.Println(data)
		switch data["code"].(type) {
		case string:
			return nil,
				errors.New(errors.RespCodeErrorCode, data["msg"].(string))
		case float64:
			if int(data["code"].(float64)) != 0 {
				return nil,
					errors.New(errors.RespCodeErrorCode, data["msg"].(string))
			}
		}
		r := data["data"].(map[string]interface{})
		return &r, nil
	})
	err = service.CheckUserModeAndFetchUserInfo()
	if err != nil {
		log.Println(err.Error())
		return
	}
	config.WaitGroup.Add(1)
	go service.MangaSignIn()
	config.WaitGroup.Wait()
	log.Println("输出错误信息：")
	for k, v := range config.ErrorSlice {
		log.Println(k, " : ", v.Error())
	}
	log.Println("全部完成")
}
