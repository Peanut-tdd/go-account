package bill

import (
	"account_check/app/console/command"
	"account_check/app/service/fapp/alipay"
	"github.com/gin-gonic/gin"
)

func FalipayBill(ctx *gin.Context) {

	billDate := ctx.Query("bill_date")
	projectId := ctx.Query("project_id")
	platformId := ctx.Query("platform_id")
	channel := ctx.Query("pay_channel")
	if projectId == "" || platformId == "" || channel == "" {
		ctx.JSONP(200,"缺少参数")
	}

	billDate = command.GetBillDate(billDate, 4)
	payConfig := command.GetConfigByQueryParmas(projectId, platformId, channel)
	alipay.BillQueryDownload(payConfig, billDate)
	ctx.JSONP(200, "success")
}
