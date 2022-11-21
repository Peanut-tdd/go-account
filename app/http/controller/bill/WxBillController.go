package bill

import (
	"account_check/app/console/command"
	"account_check/app/service/wechat"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func WxBill(c *gin.Context) {

	//入参
	billDate := c.Query("bill_date")
	projectId := c.Query("project_id")
	platformId := c.Query("platform_id")
	channel := c.Query("pay_channel")
	if projectId == "" || platformId == "" || channel == "" {
		c.JSONP(200, "缺少参数")
	}

	billDate = command.GetBillDate(billDate, 3)

	//支付参数
	payConfig := command.GetConfigByQueryParmas(projectId, platformId, channel)

	var request = make(map[string]string)
	request["appid"] = payConfig.AppId
	request["mch_id"] = payConfig.MchId
	request["nonce_str"] = string(rand.Intn(99999))
	request["sign_type"] = "MD5"
	request["bill_date"] = billDate
	request["bill_type"] = "SUCCESS"
	//request["tar_type"] = "GZIP"


	fmt.Println(request)
	wechat.GetBills(payConfig, request)
	c.JSONP(200, "success")
}
