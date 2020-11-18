package service

import (
	"GoBilibiliExpHelper/apis"
	"GoBilibiliExpHelper/config"
	"GoBilibiliExpHelper/errors"
	"GoBilibiliExpHelper/http"
)

func CheckUserModeAndFetchUserInfo() error {
	if config.AppConfig.User.Mode == "login" {
		apis.Login()
	} else if config.AppConfig.User.Mode != "token" {
		return errors.ConfigError
	}
	http.SetCookie(config.AppConfig)
	var err error
	config.AppUser, err = apis.UserInfo()
	return err
}
