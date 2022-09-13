package wechat

import (
	"account_check/app/utils"
	"log"
	"strings"
)

//GetBills 获得账单信息
func GetBills(request map[string]string) {
	sign := utils.MD5Params(request, "636x44f9y3OqN65DRbrh4Zqydobt6MBW", nil)
	request["sign"] = sign
	//map转xml
	params := toXml(request)
	//http请求
	res, _ := utils.HttpSendXmlResJson("https://api.mch.weixin.qq.com/pay/downloadbill", "post", params, nil, "")
	resArr := strings.Split(res.String(), "\r\n")
	for _, v := range resArr {
		log.Print("=====" + v + "======")
	}
	
}

//toXml map转xml
func toXml(params map[string]string) string {
	xmlParam := "<xml>"
	for key, value := range params {
		xmlParam += "<" + key + ">" + value + "</" + key + ">"
	}
	xmlParam += "</xml>"

	return xmlParam
}
