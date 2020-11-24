package apis

import (
	"GoBilibiliExpHelper/http"
)

// MangaCheckIn 漫画签到操作
func MangaCheckIn(platform string) error {
	_, err := http.POSTFORM(MangaCheckInURL,
		nil, map[string]string{
			"platform": platform,
		})
	if err != nil {
		return err
	}
	return nil
}

// MangaVipReward 获取大会员福利操作
func MangaVipReward(reasonId int) (int, error) {
	res, err := http.POST(MangaVipRewardURL,
		nil, map[string]int{
			"reason_id": reasonId,
		})
	if err != nil {
		return 0, err
	}
	return int((*res)["amount"].(float64)), nil
}
