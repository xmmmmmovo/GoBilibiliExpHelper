package config

import (
	"GoBilibiliExpHelper/models"
	"sync"
)

// AppConfig 全局设置变量
var AppConfig = new(Config)
// AppUser 全局用户变量
var AppUser *models.UserInfo
// WaitGroup 全局等待锁
var WaitGroup sync.WaitGroup
// ErrorSlice 错误信息变量
var ErrorSlice = make([]error, 0)
// ErrorSliceMutex 读写锁
var ErrorSliceMutex sync.RWMutex