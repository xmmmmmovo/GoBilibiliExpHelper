package service

import (
	"GoBilibiliExpHelper/apis"
	"GoBilibiliExpHelper/config"
	"GoBilibiliExpHelper/utils"
	"log"
)

func LiveCheckIn() {
	defer config.WaitGroup.Done()
	utils.PrintStart("直播签到")
	err := apis.LiveCheckIn()
	if err != nil {
		if err.Error() != "{\"Code\":40001,\"Msg\":\"今日已签到过,无法重复签到\"}" {
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
