package config

import (
	"GoBilibiliExpHelper/errors"
	"GoBilibiliExpHelper/utils"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
)

// LoadConfig 读取配置
func LoadConfig(path string) error {
	utils.PrintStart("读取配置")
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("打开配置文件失败！")
		return err
	}
	err = yaml.Unmarshal(fileBytes, AppConfig)
	if err != nil {
		log.Panicln("读取配置文件失败！")
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
		log.Panicln("重构配置文件失败！")
		return err
	}
	err = ioutil.WriteFile(path, res, os.ModePerm)
	if err != nil {
		log.Panicln("写入配置文件失败！")
		return err
	}
	log.Println("写入成功！")

	utils.PrintEnd("写入配置")
	return nil
}

// EnvReader 读取环境
func EnvReader() error {
	BILI_JCT := os.Getenv("BILI_JCT")
	DEDEUSERID := os.Getenv("DEDEUSERID")
	SESSDATA := os.Getenv("SESSDATA")

	if AppConfig.User.Mode == "token" {
		if AppConfig.User.Token.BILIJCT == "" {
			if BILI_JCT == "" {
				return errors.ConfigError
			} else {
				AppConfig.User.Token.BILIJCT = BILI_JCT
			}
		}

		if AppConfig.User.Token.DEDEUSERID == "" {
			if DEDEUSERID == "" {
				return errors.ConfigError
			} else {
				AppConfig.User.Token.DEDEUSERID = DEDEUSERID
			}
		}

		if AppConfig.User.Token.SESSDATA == "" {
			if SESSDATA == "" {
				return errors.ConfigError
			} else {
				AppConfig.User.Token.SESSDATA = SESSDATA
			}
		}
	}

	return nil
}
