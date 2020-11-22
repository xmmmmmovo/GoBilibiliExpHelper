package http

import (
	"bytes"
	"github.com/levigross/grequests"
	"mime/multipart"
)

// GET 请求
func GET(url string, params map[string]string, json interface{}) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
		Params:  params,
		JSON:    &json,
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
func POST(url string, params map[string]string, json interface{}) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
		Params:  params,
		JSON:    &json,
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

// POSTFORM 请求FormData
func POSTFORM(url string, params map[string]string, formData map[string]string) (*map[string]interface{}, error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range formData {
		w.WriteField(k, v)
	}
	w.Close()

	formHeader := map[string]string{
		"Content-Type": w.FormDataContentType(),
	}

	for k, v := range headers {
		formHeader[k] = v
	}

	ro := &grequests.RequestOptions{
		Cookies:     cookies,
		Headers:     formHeader,
		Params:      params,
		RequestBody: body,
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
func PUT(url string, params map[string]string, json interface{}) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
		Params:  params,
		JSON:    &json,
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
func PATCH(url string, params map[string]string, json interface{}) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
		Params:  params,
		JSON:    &json,
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
func DELETE(url string, params map[string]string, json interface{}) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
		Params:  params,
		JSON:    &json,
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
func OPTIONS(url string, params map[string]string, json interface{}) (*map[string]interface{}, error) {
	ro := &grequests.RequestOptions{
		Cookies: cookies,
		Headers: headers,
		Params:  params,
		JSON:    &json,
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
