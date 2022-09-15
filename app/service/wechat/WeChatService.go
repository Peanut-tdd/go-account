package wechat

import (
	"account_check/app/utils"
	"encoding/csv"
	"log"
	"os"
)

//GetBills 获得账单信息
func GetBills(request map[string]string) {
	sign := utils.MD5Params(request, "636x44f9y3OqN65DRbrh4Zqydobt6MBW", nil, "WECHAT")
	request["sign"] = sign
	//map转xml
	params := toXml(request)
	//http请求
	res, _ := utils.HttpSendBodyResDownLoad("https://api.mch.weixin.qq.com/pay/downloadbill", "post", params, nil, nil, "./test.csv", "")
	log.Print(res.String())
	//resArr := strings.Split(res.String(), "\r\n")
	//for _, v := range resArr {
	//	log.Print("=====" + v + "======")
	//}

}

func ReadCsv(filepath string) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		log.Println("csv文件打开失败！")
	}
	defer csvFile.Close()
	//创建csv读取接口实例
	reader := csv.NewReader(csvFile)

	//csvData, err := reader.Read()
	b, err := reader.ReadAll()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	////var oneRecord Employee
	////var allRecords []Employee
	//log.Println(csvData)
	log.Println(b)
	//for _, each := range csvData {
	//	log.Println(each)
	//}
	//os.Exit(1)
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
