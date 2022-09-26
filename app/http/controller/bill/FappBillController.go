package bill

import (
	"account_check/app/console/command"
	"account_check/app/service/fapp/alipay"
	"github.com/gin-gonic/gin"
)

func FalipayBill(ctx *gin.Context) {

	billDate := ctx.Query("bill_date")
	billDate = command.GetBillDate(billDate, 4)

	alipay.BillQueryDownload(billDate)
	ctx.JSONP(200, "success")
}
