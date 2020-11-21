package http

import "github.com/levigross/grequests"
// GET 请求
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

// POST 请求
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

// PUT 请求
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

// PATCH 请求
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

// DELETE 请求
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

// OPTIONS 请求
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
