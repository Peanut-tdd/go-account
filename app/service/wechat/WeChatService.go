package wechat

import (
	"account_check/app/model"
	"account_check/app/utils"
	"account_check/bootstrap/driver"
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	length := len(content) - 1

	if length < 0 {
		return
	}
	lencsv := len(content)

	bills := []map[string]interface{}{}

	sliceTradeNo := make([]string, lencsv)

	for index, item := range content {
		if index < 1 || index > lencsv-3 {
			continue
		}

		sliceItem := strings.Split(strings.Replace(item[0], "`", "", -1), ",")
		amount, _ := strconv.ParseFloat(sliceItem[12], 64)
		amount = amount * 100

		key, _ := strconv.Atoi(sliceItem[6])
		fmt.Println(key)
		return
		bills[key]["Number"] = sliceItem[5]
		bills[key]["TradeNo"] = sliceItem[6]
		bills[key]["TradeAt"] = utils.StringToTime(sliceItem[0])
		bills[key]["Amount"] = int(amount)
		bills[key]["PlatformId"] = 3
		bills[key]["Number"] = sliceItem[5]

		fmt.Println(bills)
		return

		sliceTradeNo = append(sliceTradeNo, sliceItem[6])

	}

	if len(sliceTradeNo) == 0 {
		return
	}

	trade_no := make([]string, lencsv)
	driver.GVA_DB.Model(&model.OrderBill{}).Pluck("trade_no", &trade_no)

	fmt.Println(trade_no)
	return

	driver.GVA_DB.Create(&bills)

	return
	//for _, b := range bills {
	//	fmt.Println(b.ID)
	//}

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
