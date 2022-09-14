package routes

import (
	"account_check/routes/RouterGroup"
	"github.com/gin-gonic/gin"
)

func RouterRegister(app *gin.Engine) {
	//log.Println("运行自定义注册路由文件 >>> ")
	RouterGroup.ApiTest(app)
}
