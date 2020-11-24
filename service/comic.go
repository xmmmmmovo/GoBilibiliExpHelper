package service

import (
	"GoBilibiliExpHelper/apis"
	"GoBilibiliExpHelper/config"
	"GoBilibiliExpHelper/errors"
	"GoBilibiliExpHelper/utils"
	"encoding/json"
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

// MangaVipReward 获取大会员福利
func MangaVipReward() {
	utils.PrintStart("领取漫画大会员福利")
	defer config.WaitGroup.Done()
	list, err := apis.GetVipPrivilegeList()
	if err != nil {
		log.Println("漫画Vip领取出现问题")
		config.ErrorSliceMutex.Lock()
		config.ErrorSlice = append(config.ErrorSlice, err)
		config.ErrorSliceMutex.Unlock()
	}
	for _, v := range *list {
		amount, err := apis.MangaVipReward(v.Type)
		if err != nil {
			e := errors.Err{}
			err = json.Unmarshal([]byte(err.Error()), &e)
			if e.Code != errors.HasReceviedRewardManga {
				config.ErrorSliceMutex.Lock()
				config.ErrorSlice = append(config.ErrorSlice, err)
				config.ErrorSliceMutex.Unlock()
			} else {
				log.Println("不能重复领取大会员福利！")
			}
			continue
		}
		log.Println("已领取：", amount, "会员券")
	}
	utils.PrintEnd("领取漫画大会员福利")
}
