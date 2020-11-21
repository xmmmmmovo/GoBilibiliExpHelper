package http

import (
	"GoBilibiliExpHelper/config"
	"net/http"
)

// cookies cookies
var cookies = make([]*http.Cookie, 3)
// headers headers
var headers = map[string]string{
	"Referer":    "https://www.bilibili.com/",
	"Connection": "keep-alive",
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70",
}
// onResponse 响应回调
var onResponse = func(data *map[string]interface{}) (*map[string]interface{}, error) {
	return data, nil
}

// SetCookie 设置cookie
func SetCookie(config *config.Config) {
	cookies[0] = &http.Cookie{
		Name:  "bili_jct",
		Value: config.User.Token.BILIJCT,
	}
	cookies[1] = &http.Cookie{
		Name:  "DedeUserID",
		Value: config.User.Token.DEDEUSERID,
	}
	cookies[2] = &http.Cookie{
		Name:  "SESSDATA",
		Value: config.User.Token.SESSDATA,
	}
}

// OnResponse 设置相应回调
func OnResponse(f func(*map[string]interface{}) (*map[string]interface{}, error)) {
	onResponse = f
}
