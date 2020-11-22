package apis

import (
	"GoBilibiliExpHelper/http"
	"log"
)

// MangaSignIn 漫画签到操作
func MangaSignIn(platform string) {
	res, err := http.POSTFORM(MangaSignInURL,
		nil, map[string]string{
			"platform": platform,
		})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res)
}
