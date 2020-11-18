package utils

import "log"

func PrintStart(tag string) {
	log.Println("***************", tag, "开始***************")
}

func PrintEnd(tag string) {
	log.Println("***************", tag, "结束***************")
}
