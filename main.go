package main

import (
	"account_check/bootstrap"

	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	// 启动服务
	bootstrap.App(HttpServer)

	return
}
