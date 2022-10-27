package RouterGroup

import (
	"account_check/app/http/Middlewares"
	"account_check/app/http/controller"
	"account_check/app/http/controller/bill"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiTest(route *gin.Engine) {
	api := route.Group("/api/", Middlewares.HttpCorsApi)

	gen := api.Group("/gen1/")
	gen.GET("1", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusNotFound, gin.H{
			"state": 200,
			"msg":   "gin1",
			"content": map[string]interface{}{
				"time": "111",
			},
		})
	})

	gen.GET("test", controller.Zzzz)

	gen.GET("wx", bill.WxBill)

	gen.GET("ks", bill.KsBill)

	gen.GET("alipay", bill.FalipayBill)

	gen.GET("delete_dir", controller.DeleteBillDir)

	gen.GET("check_coin", controller.CheckCoin)

	gen.GET("pay_config", controller.PayConfig)
}
