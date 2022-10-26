package console

import (
	"account_check/app/console/command"

	"github.com/robfig/cron/v3"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)

	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func InitConsole() {
	c := newWithSeconds()

	//todo 此处添加对应的定时任务方法
	//c.AddFunc("*/5 * * * * ?", command.Test)
	c.AddFunc("0 0 14 * * ?", command.PayCompare)       //支付账单对账
	c.AddFunc("0 0 14 * * ?", command.CoinCheckMessage) //账户虚拟币对账
	c.AddFunc("* * 23 1 * ?", command.DeleteBillDir)    //删除csv下载目录

	c.Start()
	//加载会阻塞
	//select {}
}
