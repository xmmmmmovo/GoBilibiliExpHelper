package config

import (
	"GoBilibiliExpHelper/errors"
	"GoBilibiliExpHelper/utils"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// LoadConfig 读取配置
func LoadConfig(path string) error {
	utils.PrintStart("读取配置")
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("打开配置文件失败！")
		return err
	}
	err = yaml.Unmarshal(fileBytes, AppConfig)
	if err != nil {
		log.Println("读取配置文件失败！")
		return err
	}
	log.Println("读取成功！")
	utils.PrintEnd("读取配置")
	return nil
}

// SaveConfig 写入配置
func SaveConfig(path string) error {
	utils.PrintStart("写入配置")

	res, err := yaml.Marshal(AppConfig)
	if err != nil {
		log.Println("重构配置文件失败！")
		return err
	}
	err = ioutil.WriteFile(path, res, os.ModePerm)
	if err != nil {
		log.Println("写入配置文件失败！")
		return err
	}
	log.Println("写入成功！")
	utils.PrintEnd("写入配置")
	return nil
}

// EnvReader 读取环境
func EnvReader() error {
	BiliJct := os.Getenv("BILI_JCT")
	DEDEUSERID := os.Getenv("DEDEUSERID")
	SESSDATA := os.Getenv("SESSDATA")

	if AppConfig.User.Mode == "token" {

		if AppConfig.User.Token.BILIJCT == "" {
			if BiliJct == "" {
				return errors.ConfigError
			}
			AppConfig.User.Token.BILIJCT = BiliJct
		}

		if AppConfig.User.Token.DEDEUSERID == "" {
			if DEDEUSERID == "" {
				return errors.ConfigError
			}
			AppConfig.User.Token.DEDEUSERID = DEDEUSERID
		}

		if AppConfig.User.Token.SESSDATA == "" {
			if SESSDATA == "" {
				return errors.ConfigError
			}
			AppConfig.User.Token.SESSDATA = SESSDATA
		}
	}

	return nil
}

// CheckVipNeedRun 检测大会员是否运行
func CheckVipNeedRun() bool {
	return AppUser.VipStatus == 1 &&
		AppUser.VipType == 2 &&
		AppConfig.Comic.Day == time.Now().Day()
}
