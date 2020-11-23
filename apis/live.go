package apis

import (
	"GoBilibiliExpHelper/http"
)

func LiveCheckIn() error {
	_, err := http.GET(LiveCheckInURL, nil, nil)
	return err
}
