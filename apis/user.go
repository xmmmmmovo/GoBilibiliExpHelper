package apis

import (
	"GoBilibiliExpHelper/http"
	"GoBilibiliExpHelper/models"
)

// Login 登陆操作
func Login() {
}

// UserInfo 获取用户信息操作
func UserInfo() (*models.UserInfo, error) {
	resp, err := http.GET(UserInfoURL, nil, nil)
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

// GetCoinNum 获取硬币数量操作
func GetCoinNum() (float64, error) {
	res, err := http.GET(GetCoinURL, nil, nil)
	if err != nil {
		return 0, err
	}
	return (*res)["money"].(float64), nil
}

// GetVipPrivilege 获取大会员权限列表
func GetVipPrivilegeList() (*[]models.VipPrivilege, error) {
	res, err := http.GET(MyVipPrivilegeURL, nil, nil)
	if err != nil {
		return nil, err
	}
	list := make([]models.VipPrivilege, len((*res)["list"].([]interface{})))
	for k, v := range (*res)["list"].([]interface{}) {
		vv := v.(map[string]interface{})
		list[k] = models.VipPrivilege{
			Type:       int(vv["type"].(float64)),
			State:      int(vv["state"].(float64)),
			ExpireTime: int(vv["expire_time"].(float64)),
		}
	}
	return &list, err
}
