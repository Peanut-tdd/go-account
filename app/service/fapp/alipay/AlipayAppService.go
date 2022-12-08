package alipay

import (
	"account_check/app/model"
	"account_check/app/utils"
	"account_check/bootstrap/driver"
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/smartwalle/alipay"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	client *alipay.AliPay
)

func BillQueryDownload(payConfig model.ProjectAppConfig, billDate string) bool {

	//请求
	appID := payConfig.AppId
	aliPublicKey := payConfig.AlipayPublicKey
	privateKey := payConfig.PrivateKey

	client := alipay.New(appID, aliPublicKey, privateKey, true)
	pay := alipay.BillDownloadURLQuery{}
	pay.BillDate = billDate
	pay.BillType = "trade"
	res, err := client.BillDownloadURLQuery(pay)
	if err != nil {
		fmt.Println(err)
	}

	//下载解压
	url := res.AliPayDataServiceBillDownloadURLQueryResponse.BillDownloadUrl
	if url == "" {
		return false
	}

	filePath := utils.CsvFilePath(payConfig.ProjectId, 4, 2)
	utils.UrlDownLoad(url, filePath)
	utils.UnzipFile(filePath, utils.CsvFileDir(payConfig.ProjectId, 4, 2), "\u6c47\u603b")

	//读文件

	//获取文件夹下所有的文件
	files, _ := utils.TPFuncReadDirFiles(filepath.Dir(filePath))
	for _, file := range files {
		//获得当前的csv文件
		if path.Ext(file) == ".csv" {
			//todo 读取csv文件并插入数据库中
			sliceBills := readBillCsv(file)
			saveBillData(payConfig.ProjectId, sliceBills)

		}
	}
	return true

}

func readBillCsv(filepath string) [][]string {

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

	return content
}

//GBK转utf8的方法
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func saveBillData(projectId uint, content [][]string) {
	//csv长度判断
	length := len(content) - 1
	if length < 0 {
		return
	}

	lencsv := len(content)

	mapBills := make(map[string]map[string]interface{})
	sliceTradeNo := make([]string, 0)

	for idx, item := range content {
		if idx < 5 || idx > lencsv-5 {
			continue
		}

		sliceItem := strings.Split(strings.Replace(item[0], "\t", "", -1), ",")

		//for k, v := range sliceItem {
		//	fmt.Println("k:%v,v:%v", k, v)
		//}
		//return

		amount, _ := strconv.ParseFloat(sliceItem[11], 64)
		if int(amount*100) <= 1 {
			continue
		}

		mapBills[sliceItem[0]] = map[string]interface{}{
			"ProjectId":  projectId,
			"Number":     sliceItem[1],
			"TradeNo":    sliceItem[0],
			"TradeAt":    sliceItem[5],
			"Amount":     int(amount * 100),
			"PlatformId": 4,
			"PayChannel": 2,
			"CreatedAt":  utils.CurrentDateTime(),
			"UpdatedAt":  utils.CurrentDateTime(),
		}
		sliceTradeNo = append(sliceTradeNo, sliceItem[0])
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
