package http

import "github.com/levigross/grequests"

func GET(url string) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
	}
	res, err := grequests.Get(url, ro)
	if err != nil {
		return nil, err
	}
	resp := new(map[string]interface{})
	err = res.JSON(resp)
	if err != nil {
		return nil, err
	}
	return onResponse(resp)
}

func POST(url string) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
	}
	res, err := grequests.Post(url, ro)
	if err != nil {
		return nil, err
	}
	resp := new(map[string]interface{})
	err = res.JSON(resp)
	if err != nil {
		return nil, err
	}
	return onResponse(resp)
}

func PUT(url string) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
	}
	res, err := grequests.Put(url, ro)
	if err != nil {
		return nil, err
	}
	resp := new(map[string]interface{})
	err = res.JSON(resp)
	if err != nil {
		return nil, err
	}
	return onResponse(resp)
}

func PATCH(url string) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
	}
	res, err := grequests.Patch(url, ro)
	if err != nil {
		return nil, err
	}
	resp := new(map[string]interface{})
	err = res.JSON(resp)
	if err != nil {
		return nil, err
	}
	return onResponse(resp)
}

func DELETE(url string) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
	}
	res, err := grequests.Delete(url, ro)
	if err != nil {
		return nil, err
	}
	resp := new(map[string]interface{})
	err = res.JSON(resp)
	if err != nil {
		return nil, err
	}
	return onResponse(resp)
}

func OPTIONS(url string) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
	}
	res, err := grequests.Options(url, ro)
	if err != nil {
		return nil, err
	}
	resp := new(map[string]interface{})
	err = res.JSON(resp)
	if err != nil {
		return nil, err
	}
	return onResponse(resp)
}
