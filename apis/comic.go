package apis

import (
	"GoBilibiliExpHelper/http"
	"log"
)

func MangaSignIn(platform string) {
	res, err := http.POST(MangaSignInURL,
		map[string]string{
			"platform": platform,
		}, nil)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res)
}
