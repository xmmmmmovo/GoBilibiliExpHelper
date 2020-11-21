package config

import (
	"GoBilibiliExpHelper/models"
	"sync"
)

// 全局设置变量
var AppConfig = new(Config)
var AppUser *models.UserInfo
var WaitGroup sync.WaitGroup