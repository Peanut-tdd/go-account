package bootstrap

import (
	"account_check/bootstrap/driver"
	"account_check/routes"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func App(HttpServer *gin.Engine) {
	// Gin服务
	HttpServer = gin.Default()

	// //加载配置文件
	// driver.InitConfig()
	// //加载gorm
	// driver.InitGorm()
	// 捕捉接口运行耗时（必须排第一）

	// 设置全局ctx参数（必须排第二）

	// 拦截应用500报错，使之可视化

	// Gin运行时：release、debug、test

	// 注册必要路由，处理默认路由、静态文件路由、404路由等

	// 注册其他路由，可以自定义
	routes.RouterRegister(HttpServer)
	// 初始化定时器（立即运行定时器）

	// 实际访问网址和端口
	host := driver.AllConfig.Server.Host + ":" + strconv.Itoa(driver.AllConfig.Server.Port)

	// 终端提示
	err := HttpServer.Run(host)
	if err != nil {
		log.Println("http服务遇到错误a。运行中断,error:", err.Error())
		log.Println("提示：注意端口被占时应该首先更改对外暴露的端口，而不是微服务的端口。")
		os.Exit(200)
	}
}
