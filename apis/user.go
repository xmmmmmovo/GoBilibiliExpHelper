package apis

import (
	"GoBilibiliExpHelper/http"
	"GoBilibiliExpHelper/models"
)

func Login() {
}

func UserInfo() (*models.UserInfo, error) {
	resp, err := http.GET(USER_INFO_URL)
	if err != nil {
		return nil, err
	}
	data := *resp
	return &models.UserInfo{
		IsLogin:   data["isLogin"].(bool),
		Uname:     data["uname"].(string),
		VipStatus: int(data["vipStatus"].(float64)),
		VipType:   int(data["vipType"].(float64)),
		CurrentLevel: int(data["level_info"].
		(map[string]interface{})["current_level"].(float64)),
		Money: data["money"].(float64),
		BcoinBalance: data["wallet"].
		(map[string]interface{})["bcoin_balance"].(float64),
		CouponBalance: data["wallet"].
		(map[string]interface{})["coupon_balance"].(float64),
	}, nil
}

func GetCoinNum() float64 {
	return 0
}
