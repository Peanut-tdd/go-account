package command

import (
	"account_check/app/model"
	"account_check/app/service/dingding"
	"account_check/app/service/fapp/alipay"
	"account_check/app/service/kuaishou"
	"account_check/app/service/wechat"
	"account_check/app/utils"
	"account_check/bootstrap/driver"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func PayCompare() {

	var message string
	//快手账单message
	ksBillDiffNumbers, ksOrderDiffNumbers := KsPayCompare()
	ksCount := len(ksBillDiffNumbers) + len(ksOrderDiffNumbers)
	if ksCount > 0 {
		message += "\n\n--------------------\n\n快手-账单异常，异常数：" + strconv.Itoa(ksCount)
	}

	//todo 微信账单message
	wxBillDiffNumber, wxOrderDiffNumber := wxPayBillCompare("", 3)
	wxDiffCount := len(wxBillDiffNumber) + len(wxOrderDiffNumber)
	if wxDiffCount > 0 {
		message += "\n\n--------------------\n\n微信账单异常，异常数：" + strconv.Itoa(wxDiffCount)
	}

	//todo 支付宝账单
	fappAlipayBillDiffNumber, fappAlipayOrderDiffNumber := fappAliPayBillCompare("", 4)
	fappAlipayDiffCount := len(fappAlipayBillDiffNumber) + len(fappAlipayOrderDiffNumber)
	if fappAlipayDiffCount > 0 {
		message += "\n\n--------------------\n\n快应用支付宝账单异常，异常数：" + strconv.Itoa(fappAlipayDiffCount)
	}
	title := "充值对账异常账单提醒"
	if message != "" {
		sendMessage(title+message, title)
	}
}

//KsPayCompare 快手比较
func KsPayCompare() ([]string, []string) {

	//获得前一天时间
	//currentTime := time.Now()
	currentTime, _ := time.Parse("2006-01-02 15:04:05", "2022-08-20 00:00:00")

	yesTime := currentTime.AddDate(0, 0, -1)
	currentTimeFormat := currentTime.Format("20060102000000")
	yesTimeFormat := yesTime.Format("20060102000000")
	syncKsBill(yesTimeFormat, currentTimeFormat)
	//处理时间
	startTime := yesTime.Format("2006-01-02 00:00:00")
	endTime := currentTime.Format("2006-01-02 00:00:00")
	billNumbers := getBillNumbers(2, 0, startTime, endTime)
	orderNumbers := getOrderNumbers(2, 0, startTime, endTime)
	//订单号对比，如果存在差异，发送钉钉消息
	billDiffNumbers, orderDiffNumbers := utils.Arrcmp(billNumbers, orderNumbers)

	return billDiffNumbers, orderDiffNumbers

}

//syncKsBill 同步快手账单
func syncKsBill(startDate string, endData string) {
	var request = make(map[string]string)
	request["app_id"] = driver.GVA_VP.GetString("ks.app_id")
	request["start_date"] = startDate
	request["end_date"] = endData
	request["bill_type"] = "PAY"
	kuaishou.GetBills(request)
}

//1-同步前一天或指定日期订单
func getBillNumbers(platFormId int, payChannel int, startDate string, endData string) []string {
	var numbers []string

	sqlModel := driver.GVA_DB.Model(&model.OrderBill{}).
		Where("platform_id = ?", platFormId).
		Where("trade_at between ? and ?", startDate, endData)
	if payChannel > 0 {
		sqlModel = sqlModel.Where("pay_channel = ?", payChannel)
	}
	sqlModel.Pluck("number", &numbers)

	return numbers
}

//getOrderNumbers 2-查询同步日期订单sql
func getOrderNumbers(platFormId int, payChannel int, startDate string, endData string) []string {
	var numbers []string
	sqlModel := driver.GVA_DB.Model(&model.Orders{}).
		Where("platform_id = ?", platFormId).
		Where("pay_success_time between ? and ?", startDate, endData).
		Where("status = 1").
		Where("trade_no != ''")
	if payChannel > 0 {
		sqlModel = sqlModel.Where("pay_channel = ?", payChannel)
	}
	sqlModel.Pluck("number", &numbers)

	return numbers
}

//4-快手数据与order查询数据比较，并钉钉通知
func sendMessage(message string, title string) {
	dingding.SendGroup(message, "chat2ec214da47216a95e7ee73ee3760d191", title)
}

func wxPayBillCompare(compareBillDate string, PlatFormId int) (wxBillDiffNumber, wxOrderDiffNumber []string) {

	//拉取昨日账单
	yesBillDate := GetBillDate(compareBillDate, PlatFormId) //昨日账单日期

	fmt.Println(yesBillDate)
	requestParams := make(map[string]string)
	requestParams["appid"] = driver.AllConfig.Wx.AppId
	requestParams["mch_id"] = driver.AllConfig.Wx.MchId
	requestParams["nonce_str"] = string(rand.Intn(99999))
	requestParams["sign_type"] = "MD5"
	requestParams["bill_date"] = yesBillDate
	requestParams["bill_type"] = "SUCCESS"
	wechat.GetBills(requestParams)

	//对比账单
	start_date, end_date := getBillDateBetween(compareBillDate, -1)

	billNumbers := getBillNumbers(3, 1, start_date, end_date)
	orderNumbers := getOrderNumbers(3, 1, start_date, end_date)
	//订单号对比，如果存在差异，发送钉钉消息
	wxBillDiffNumber, wxOrderDiffNumber = utils.Arrcmp(billNumbers, orderNumbers)

	return

}

func fappAliPayBillCompare(compareBillDate string, PlatFormId int) (fappAlipayBillDiffNumber, fappAlipayOrderDiffNumber []string) {
	//拉取昨日账单
	yesBillDate := GetBillDate(compareBillDate, PlatFormId) //昨日账单日期
	alipay.BillQueryDownload(yesBillDate)

	//对比账单
	start_date, end_date := getBillDateBetween(compareBillDate, -1)

	billNumbers := getBillNumbers(4, 2, start_date, end_date)
	orderNumbers := getOrderNumbers(4, 2, start_date, end_date)
	//订单号对比，如果存在差异，发送钉钉消息
	fappAlipayBillDiffNumber, fappAlipayOrderDiffNumber = utils.Arrcmp(billNumbers, orderNumbers)
	return
}

//拉取某天的账单
func GetBillDate(billDate string, PlatFormId int) string {
	var queryBillDate string
	if billDate != "" {
		parseTime, _ := time.ParseInLocation("2006-01-02", billDate, time.Local)
		switch PlatFormId {
		case 3:
			queryBillDate = parseTime.Format("20060102")
			break
		case 4:
			queryBillDate = parseTime.Format("2006-01-02")
			break
		}

	} else {
		switch PlatFormId {
		case 3:
			queryBillDate = time.Now().AddDate(0, 0, -1).Format("20060102")
		case 4:
			queryBillDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02")

		}
	}
	return queryBillDate

}

//对比两个时间节点的账单
func getBillDateBetween(date string, diff int) (start_date, end_date string) {
	if date != "" {
		parseStartDate, _ := time.ParseInLocation("2006-01-02", date, time.Local)
		end_date_time := parseStartDate.AddDate(0, 0, 1)
		end_date = end_date_time.Format("2006-01-02 00:00:00")
		start_date = end_date_time.AddDate(0, 0, diff).Format("2006-01-02 00:00:00")

	} else {
		start_date = time.Now().AddDate(0, 0, diff).Format("2006-01-02 00:00:00")
		end_date = time.Now().Format("2006-01-02 00:00:00")

	}
	return

}
