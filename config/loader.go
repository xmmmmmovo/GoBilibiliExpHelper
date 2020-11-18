package config

import (
	"GoBilibiliExpHelper/utils"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
)

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
