package apis

import (
	"GoBilibiliExpHelper/http"
)

// MangaSignIn 漫画签到操作
func MangaSignIn(platform string) error {
	_, err := http.POSTFORM(MangaSignInURL,
		nil, map[string]string{
			"platform": platform,
		})
	if err != nil {
		return err
	}
	return nil
}
