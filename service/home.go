package service

import (
	"GoBilibiliExpHelper/apis"
	"GoBilibiliExpHelper/config"
	"GoBilibiliExpHelper/errors"
	"GoBilibiliExpHelper/utils"
	"encoding/json"
	"log"
)

// HomeVipReward 获取主站大会员福利
func HomeVipReward() {
	utils.PrintStart("领取主站大会员福利")
	defer config.WaitGroup.Done()
	list, err := apis.GetVipPrivilegeList()
	if err != nil {
		log.Println("主站Vip领取出现问题")
		config.ErrorSliceMutex.Lock()
		config.ErrorSlice = append(config.ErrorSlice, err)
		config.ErrorSliceMutex.Unlock()
	}
	for _, v := range *list {
		err := apis.HomeVipReward(v.Type, config.AppConfig.User.Token.BILIJCT)
		if err != nil {
			e := errors.Err{}
			err = json.Unmarshal([]byte(err.Error()), &e)
			if e.Code != errors.HasReceviedRewardHome {
				config.ErrorSliceMutex.Lock()
				config.ErrorSlice = append(config.ErrorSlice, err)
				config.ErrorSliceMutex.Unlock()
			} else {
				log.Println("不能重复领取大会员福利！")
			}
			continue
		}
		log.Println("已领取会员福利")
	}
	utils.PrintEnd("领取主站大会员福利")
}
