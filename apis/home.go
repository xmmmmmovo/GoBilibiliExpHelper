package apis

import (
	"GoBilibiliExpHelper/http"
	"strconv"
)

// HomeVipReward 获取大会员福利操作
func HomeVipReward(typeID int, csrf string) error {
	_, err := http.POSTFORM(VipPrivilegeReceiveURL,
		nil, map[string]string{
			"type": strconv.Itoa(typeID),
			"csrf": csrf,
		})
	if err != nil {
		return err
	}
	return nil
}
