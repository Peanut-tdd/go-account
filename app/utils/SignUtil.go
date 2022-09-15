package utils

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func MD5Params(params map[string]string, key string, filters []string, platform string) string {
	// 将请求参数的key提取出来，并排好序
	newKeys := make([]string, 0)
	for k, _ := range params {
		//需要过滤的签名
		flag := true
		for _, filter := range filters {
			if k == filter {
				flag = false
			}
		}
		if flag {
			newKeys = append(newKeys, k)
		}
	}
	sort.Strings(newKeys)
	var originStr string
	// 将输入进行标准化的处理
	for _, v := range newKeys {
		originStr += fmt.Sprintf("%v=%v&", v, params[v])
	}
	if platform == "WECHAT" {
		originStr += fmt.Sprintf("key=%v", key)
	} else if platform == "KS" {
		originStr = strings.Trim(originStr, "&") + key
	}

	// 使用md5算法进行处理
	sign := MD5(originStr)

	return sign

}

//MD5 md5加密
func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
