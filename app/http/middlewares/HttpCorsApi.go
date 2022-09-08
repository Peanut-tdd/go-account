package Middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpCorsApi(ctx *gin.Context) {
	method := ctx.Request.Method
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//
	////
	ctx.Header("Content-type", "application/json; charset=utf-8")
	ctx.Header("Cache-Control", "max-age=0")
	// ctx.Header("X-Powered-By", "ginvel.com; "+Common.ServerInfo["go_version"])
	ctx.Header("Author", "fyonecon")
	// ctx.Header("Timezone", Common.ServerInfo["timezone"])
	// ctx.Header("Date", Common.GetTimeDate("Y-m-d H:i:s"))

	//是否允许后续请求携带认证信息,该值只能是true,否则不返回
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}

	ctx.Next()
}
