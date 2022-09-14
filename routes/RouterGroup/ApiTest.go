package RouterGroup

import (
	"account_check/app/Http/Middlewares"
	"account_check/app/Http/controller"
	"account_check/app/http/controller/bill"
	"net/http"

	"github.com/gin-gonic/gin"
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
}
