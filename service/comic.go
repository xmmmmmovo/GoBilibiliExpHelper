package service

import (
	"GoBilibiliExpHelper/apis"
	"GoBilibiliExpHelper/config"
	"GoBilibiliExpHelper/utils"
	"log"
)

// MangaCheckIn 漫画签到
func MangaCheckIn() {
	utils.PrintStart("漫画签到")
	defer config.WaitGroup.Done()
	err := apis.MangaCheckIn(config.AppConfig.Comic.Platform)
	if err != nil {
		log.Println("您哔哩哔哩漫画已签过到了！")
	} else {
		log.Println("您哔哩哔哩漫画已成功签到！")
	}
	utils.PrintEnd("漫画签到")
}
