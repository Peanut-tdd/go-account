package main

import (
	"account_check/bootstrap"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	//压缩文件解压测试
	//err := utils.Unzip("test.csv.zip", ".")
	//fmt.Println("err is", err)
	//return
	//-------------------快手拉去账单测试
	//var request = make(map[string]string)
	//request["app_id"] = "ks695806146341101215"
	//request["start_date"] = "20220819000000"
	//request["end_date"] = "20220820000000"
	//request["bill_type"] = "PAY"
	//kuaishou.GetBills(request)

	//-------------------微信拉去账单测试
	//var request = make(map[string]string)
	//request["appid"] = "wxe814fe9772dcc6df"
	//request["mch_id"] = "1616436352"
	//request["nonce_str"] = string(rand.Intn(99999))
	//request["sign_type"] = "MD5"
	//request["bill_date"] = "20220819"
	//request["bill_type"] = "SUCCESS"
	////request["tar_type"] = "GZIP"
	//
	//wechat.GetBills(request)
	// 启动服务
	bootstrap.App(HttpServer)

	return
}
