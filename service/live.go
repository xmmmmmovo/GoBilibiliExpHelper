package service

import (
	"GoBilibiliExpHelper/apis"
	"GoBilibiliExpHelper/config"
	"GoBilibiliExpHelper/errors"
	"GoBilibiliExpHelper/utils"
	"encoding/json"
	"log"
)

// LiveCheckIn 直播签到
func LiveCheckIn() {
	defer config.WaitGroup.Done()
	utils.PrintStart("直播签到")
	err := apis.LiveCheckIn()
	if err != nil {
		e := errors.Err{}
		err = json.Unmarshal([]byte(err.Error()), &e)
		if e.Code == errors.RespCodeErrorCode {
			config.ErrorSliceMutex.Lock()
			config.ErrorSlice = append(config.ErrorSlice, err)
			config.ErrorSliceMutex.Unlock()
		} else {
			log.Println("您哔哩哔哩直播已签过到了！")
		}
	} else {
		log.Println("您哔哩哔哩直播已签过到了！")
	}
	utils.PrintEnd("直播签到")
}
