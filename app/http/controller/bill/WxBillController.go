package bill

import (
	"account_check/app/console/command"
	"account_check/app/service/wechat"
	"account_check/bootstrap/driver"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func WxBill(c *gin.Context) {
	billDate := c.Query("bill_date")
	billDate =command.GetBillDate(billDate,3)


	var request = make(map[string]string)
	request["appid"] = driver.AllConfig.Wx.AppId
	request["mch_id"] = driver.AllConfig.Wx.MchId
	request["nonce_str"] = string(rand.Intn(99999))
	request["sign_type"] = "MD5"
	request["bill_date"] = billDate
	request["bill_type"] = "SUCCESS"
	//request["tar_type"] = "GZIP"
	wechat.GetBills(request)
	c.JSONP(200, "success")
}
