package bootstrap

import (
	"account_check/app/console"
	"account_check/bootstrap/driver"
)

func init() {
	//加载配置文件
	driver.InitConfig()
	//加载gorm
	driver.InitGorm()
	//加载redis
	driver.InitRedis()

	console.InitConsole()

}
