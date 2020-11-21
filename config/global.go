package config

import (
	"GoBilibiliExpHelper/models"
	"sync"
)

var AppConfig = new(Config)
var AppUser *models.UserInfo
var WaitGroup sync.WaitGroup