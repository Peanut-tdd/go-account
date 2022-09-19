package wechat

import (
	"account_check/app/utils"
	"account_check/bootstrap/driver"
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

//GetBills 获得账单信息
func GetBills(request map[string]string) {
	//ReadCsv("./test.csv")
	//return
	sign := utils.MD5Params(request, driver.AllConfig.Wx.Key, nil, "WECHAT")
	request["sign"] = sign

	//fmt.Println(request)
	//map转xml
	params := toXml(request)
	//http请求
	filepath := utils.CsvFilePath(1)
	res, err := utils.HttpSendBodyResDownLoad("https://api.mch.weixin.qq.com/pay/downloadbill", "post", params, nil, nil, filepath, "")

	if err != nil {
		log.Fatal("获取csv文件信息失败，err is %+v", err)
	}
	log.Println(res.String())
	ReadCsv(filepath)
	//resArr := strings.Split(res.String(), "\r\n")
	//for _, v := range resArr {
	//	log.Print("=====" + v + "======")
	//}

}

//func wxBillFix(res string) {
//	resArr := strings.Split(strings.Replace(res, ",", " ", -1), "`")
//	resLen := (len(resArr) - 6) / 24
//
//	for k, v := range resArr {
//		fmt.Println("K:%v，v:%v", k, v)
//	}
//
//	return
//
//	fmt.Println(resArr)
//	//return
//	sliceMap := make([]map[string]interface{}, 0)
//	var index int
//	for i := 0; i < resLen; i++ {
//		index = 24 * i
//
//		sliceMap[index+7]["wechat_order_no"] = resArr[i+6]
//		sliceMap[i]["amount"] = resArr[+12]
//	}
//	fmt.Println(sliceMap)
//}

func readCsvFromByte(str string) {

	dataByte := []byte(str)
	content, err := csv.NewReader(bytes.NewReader(dataByte)).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, row := range content {
		fmt.Println("k=", k)
		for _, row2 := range row {
			fmt.Println(row2)
		}
	}
}

func ReadCsv(filepath string) {
	fmt.Println(filepath)
	csvFile, err := os.Open(filepath)
	if err != nil {
		log.Println("csv文件打开失败！")
	}
	defer csvFile.Close()
	//创建csv读取接口实例
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	reader.LazyQuotes = true

	//csvData, err := reader.Read()
	content, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}

	fmt.Println(content)

	length := len(content) - 1
	if length < 0 {
		return
	}
	len := len(content)


	type Bill struct {
		Trade_No string
		Amount string
	}
	//bills := make(map[string]interface{})
	for index, item := range content {
		if index < 1 || index > len-3 {
			continue
		}



		fmt.Println("k:%v,value:%v\n", index, item)
	}

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
