package utils

import "log"

// PrintStart 打印service开始
func PrintStart(tag string) {
	log.Println("***************", tag, "开始***************")
}

// PrintEnd 打印service结束
func PrintEnd(tag string) {
	log.Println("***************", tag, "结束***************")
}
