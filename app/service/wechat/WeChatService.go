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
	"regexp"
	"strconv"
	"strings"
)

//GetBills 获得账单信息
func GetBills(payConfig model.ProjectAppConfig, request map[string]string) {
	//ReadCsv("./test.csv")
	//return
	sign := utils.MD5Params(request, payConfig.Key, nil, "WECHAT")
	request["sign"] = sign

	//map转xml
	params := toXml(request)
	//http请求
	filepath := utils.CsvFilePath(payConfig.ProjectId, 1, 1)
	//fmt.Println("filepath:", filepath)
	_, err := utils.HttpSendBodyResDownLoad("https://api.mch.weixin.qq.com/pay/downloadbill", "post", params, nil, nil, filepath, "")

	if err != nil {
		log.Fatal("获取csv文件信息失败，err is %+v", err)
	}

	ReadCsv(payConfig.ProjectId, filepath)
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

func ReadCsv(projectId uint, filepath string) {
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

	//csv长度判断
	length := len(content) - 1
	if length < 0 {
		return
	}

	lencsv := len(content)

	mapBills := make(map[string]map[string]interface{})
	sliceTradeNo := make([]string, 0)
	//测试环境用户订单需过滤
	testUserOpendIds := []string{"osJgh5do_edzSgrHNZdZ-rLUj30A", "osJgh5f1BOu0jcMmH_ZrhVoZabPI", "osJgh5ZKKB9KYOdF7cer9gMl_BQk", "osJgh5WcYdAhG4x8X4ykV3-UMCrQ", "osJgh5X5-F9UeVSpzAIaDF0Uz1uE", "osJgh5RaKq-hgZDAUUuRCQ8W5uhY", "osJgh5SPpr01LhCLFW4T6NUf_WM4", "osJgh5aSJNC4f_9MBTHE8ef-QpL0", "osJgh5SCHOQ7OPmRRH6cFREOjGis", "oqk-H5czjDZvyuIBViQ8IztgWD-M", "oqk-H5fTWR4SCTuPFzgM84ygKUGQ"}

	for index, item := range content {
		//是否入库
		goingOn := true

		re := regexp.MustCompile("<xml>")
		match := re.MatchString(item[0])
		if match {
			break
		}

		if index < 1 || index > lencsv-3 {
			continue
		}

		sliceItem := strings.Split(strings.Replace(item[0], "`", "", -1), ",")
		amount, _ := strconv.ParseFloat(sliceItem[12], 64)
		amount = amount * 100

		//过滤测试账号
		for _, openid := range testUserOpendIds {
			if openid == sliceItem[7] {
				goingOn = false
				break
			}
		}
		if goingOn == false {
			continue
		}

		key := sliceItem[5]
		mapBills[key] = map[string]interface{}{
			"ProjectId":  projectId,
			"Number":     sliceItem[6],
			"TradeNo":    sliceItem[5],
			"TradeAt":    utils.StringToTime(sliceItem[0]),
			"Amount":     int(amount),
			"PlatformId": 3,
			"CreatedAt":  utils.CurrentDateTime(),
			"UpdatedAt":  utils.CurrentDateTime(),
		}

		sliceTradeNo = append(sliceTradeNo, sliceItem[5])

	}

	if len(sliceTradeNo) == 0 {
		return
	}

	//判断数据新增还是编辑
	db_trade_no := make([]string, 0)
	driver.GVA_DB.Model(&model.OrderBill{}).Where("trade_no in ?", sliceTradeNo).Pluck("trade_no", &db_trade_no)

	insertSliceMapBills := make([]map[string]interface{}, 0)
	//updateSliceMapBills := make([]map[string]interface{}, 0)
	insertSliceMapBills = utils.CheckInsertData(db_trade_no, insertSliceMapBills, mapBills)

	if len(insertSliceMapBills) == 0 {
		return
	}

	//插入数据
	driver.GVA_DB.Model(&model.OrderBill{}).Create(insertSliceMapBills)
	return

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
