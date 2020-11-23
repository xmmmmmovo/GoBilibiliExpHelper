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
		log.Println("响应: ", data)
		switch data["code"].(type) {
		case string:
			return nil,
				errors.New(errors.RespCodeErrorCode, data["msg"].(string))
		case float64:
			if int(data["code"].(float64)) != 0 {
				eMsg := ""
				if data["msg"] != nil {
					eMsg = "msg"
				} else {
					eMsg = "message"
				}
				return nil,
					errors.New(errors.RespCodeErrorCode, data[eMsg].(string))
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
	if config.AppConfig.Comic.On {
		config.WaitGroup.Add(1)
		go service.MangaCheckIn()
	}
	if config.AppConfig.Live.On {
		config.WaitGroup.Add(1)
		go service.LiveCheckIn()
	}
	if config.AppConfig.Home.On {

	}
	config.WaitGroup.Wait()
	log.Println("输出错误信息：")
	for k, v := range config.ErrorSlice {
		log.Println(k, " : ", v.Error())
	}
	log.Println("全部完成")
}
