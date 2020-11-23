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
