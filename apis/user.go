package apis

import (
	"GoBilibiliExpHelper/http"
	"GoBilibiliExpHelper/models"
	"log"
)

func Login() {
}

func UserInfo() (*models.UserInfo, error) {
	resp, err := http.GET(USER_INFO_URL)
	if err != nil {
		return nil, err
	}
	log.Println(resp)
	return &models.UserInfo{
		IsLogin:       false,
		Uname:         "",
		VipStatus:     0,
		VipType:       0,
		CurrentLevel:  0,
		Money:         0,
		BcoinBalance:  0,
		CouponBalance: 0,
	}, nil
}

func GetCoinNum() float64 {
	return 0
}
