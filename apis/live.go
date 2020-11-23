package apis

import (
	"GoBilibiliExpHelper/http"
)

// LiveCheckIn 直播签到操作
func LiveCheckIn() error {
	_, err := http.GET(LiveCheckInURL, nil, nil)
	return err
}
