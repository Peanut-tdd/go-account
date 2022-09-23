package bill

import (
	"account_check/app/service/fapp/alipay"
	"github.com/gin-gonic/gin"
)

func FalipayBill(ctx *gin.Context) {

	alipay.BillQueryDownload()

	ctx.JSONP(200, "success")
}
