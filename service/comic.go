package service

import (
	"GoBilibiliExpHelper/apis"
	"GoBilibiliExpHelper/config"
)

// MangaSignIn 漫漶签到
func MangaSignIn() {
	defer config.WaitGroup.Done()
	apis.MangaSignIn(config.AppConfig.Comic.Platform)
}
