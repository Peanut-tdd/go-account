package utils

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

const httpClientTimeOut = 3
const httpClientRetryCount = 3

// HttpGetResJson get request and json response
func HttpGetResJson(url string, queryParams map[string]string, headers map[string]string, result interface{}) (res *resty.Response, err error) {
	client := resty.New()

	client.SetTimeout(time.Second * httpClientTimeOut)
	client.SetRetryCount(httpClientRetryCount)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	req := client.R().
		SetQueryParams(queryParams).
		SetHeader("Accept", "application/json").
		SetResult(result)
	//设置header头
	for key, value := range headers {
		req = req.SetHeader(key, value)
	}
	res, err = req.Get(url)

	return res, err
}

// HttpSendFormResJson send formData and response json
func HttpSendFormResJson(url, method string, formData map[string]string, headers map[string]string, result interface{}) (res *resty.Response, err error) {
	client := resty.New()

	client.SetTimeout(time.Second * httpClientTimeOut)
	client.SetRetryCount(httpClientRetryCount)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	req := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").
		SetFormData(formData).
		SetResult(result)
	//设置header头
	for key, value := range headers {
		req = req.SetHeader(key, value)
	}
	switch strings.ToLower(method) {
	case "post":
		res, err = req.Post(url)
	case "put":
		res, err = req.Put(url)
	case "patch":
		res, err = req.Patch(url)
	case "delete":
		res, err = req.Delete(url)
	case "options":
		res, err = req.Options(url)
	default:
		res, err = req.Head(url)
	}

	return res, err
}

func HttpSendXmlResJson(url, method string, body interface{}, headers map[string]string, result interface{}) (res *resty.Response, err error) {
	client := resty.New()

	client.SetTimeout(time.Second * httpClientTimeOut)
	client.SetRetryCount(httpClientRetryCount)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req := client.R().
		SetBody(body).
		SetResult(result)
	//设置header头
	for key, value := range headers {
		req = req.SetHeader(key, value)
	}
	switch strings.ToLower(method) {
	case "post":
		res, err = req.Post(url)
	case "put":
		res, err = req.Put(url)
	case "patch":
		res, err = req.Patch(url)
	case "delete":
		res, err = req.Delete(url)
	case "options":
		res, err = req.Options(url)
	default:
		res, err = req.Head(url)
	}

	return res, err
}

// HttpSendJsonResJson send json and response json
func HttpSendJsonResJson(url, method string, body interface{}, headers map[string]string, result interface{}) (res *resty.Response, err error) {
	client := resty.New()

	client.SetTimeout(time.Second * httpClientTimeOut)
	client.SetRetryCount(httpClientRetryCount)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	req := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(body).
		SetResult(result)
	//设置header头
	for key, value := range headers {
		req = req.SetHeader(key, value)
	}
	switch strings.ToLower(method) {
	case "post":
		res, err = req.Post(url)
	case "put":
		res, err = req.Put(url)
	case "patch":
		res, err = req.Patch(url)
	case "delete":
		res, err = req.Delete(url)
	case "options":
		res, err = req.Options(url)
	default:
		res, err = req.Head(url)
	}

	return res, err
}
