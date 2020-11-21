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
		log.Panicln(err.Error())
	}
	http.OnResponse(func(resp *map[string]interface{}) (*map[string]interface{}, error) {
		data := * resp
		if int(data["code"].(float64)) != 0 {
			return nil, errors.RespCodeError
		}
		r := data["data"].(map[string]interface{})
		return &r, nil
	})
	err = service.CheckUserModeAndFetchUserInfo()
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	config.WaitGroup.Wait()
	log.Println("全部完成")
}
