package utils

import (
	"crypto/md5"
	"fmt"
	"log"
	"sort"
)

func MD5Params(params map[string]string, key string, filters []string) string {
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
	log.Println(originStr)
	originStr += fmt.Sprintf("key=%v", key)
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