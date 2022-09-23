package command

import (
	"account_check/app/model"
	"account_check/app/service/kuaishou"
	"account_check/app/utils"
	"account_check/bootstrap/driver"
	"fmt"
	"time"
)

//KsPayCompare 快手比较
func KsPayCompare() {
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

	fmt.Println(billDiffNumbers)
	fmt.Println("------------------------")
	fmt.Println(orderDiffNumbers)
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

//3-查询order表近3天快手支付成功订单

//4-快手数据与order查询数据比较，并钉钉通知
