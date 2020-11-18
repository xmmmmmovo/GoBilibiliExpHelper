package service

import (
	"GoBilibiliExpHelper/apis"
	"GoBilibiliExpHelper/config"
	"GoBilibiliExpHelper/errors"
	"GoBilibiliExpHelper/http"
	"GoBilibiliExpHelper/utils"
	"log"
)

func CheckUserModeAndFetchUserInfo() error {
	utils.PrintStart("获取用户信息")
	if config.AppConfig.User.Mode == "login" {
		apis.Login()
	} else if config.AppConfig.User.Mode != "token" {
		return errors.ConfigError
	}
	http.SetCookie(config.AppConfig)
	var err error
	config.AppUser, err = apis.UserInfo()
	if err != nil {
		return err
	}
	log.Println("成功获取到用户信息：")
	var isLogin string
	if config.AppUser.VipStatus == 0 {
		isLogin = "未签到"
	} else {
		isLogin = "已签到"
	}
	log.Println("是否成功登录签到：", isLogin)
	log.Println("用户名：", config.AppUser.Uname)
	var isVip string
	if config.AppUser.VipStatus == 0 {
		isVip = "否"
	} else {
		isVip = "是"
	}
	log.Println("是否是大会员：", isVip)
	var vipName string
	switch config.AppUser.VipType {
	case 0:
		vipName = "无"
	case 1:
		vipName = "月度"
	case 2:
		vipName = "年度"
	}
	log.Println("大会员等级：", vipName)
	log.Println("当前等级：", config.AppUser.CurrentLevel, "级")
	log.Println("硬币数：", config.AppUser.Money, "个")
	log.Println("B币数：", config.AppUser.BcoinBalance, "个")
	log.Println("每月奖励B币数：", config.AppUser.CouponBalance, "个")

	utils.PrintEnd("获取用户信息")
	return err
}
