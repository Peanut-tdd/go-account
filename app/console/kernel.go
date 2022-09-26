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
	c.AddFunc("*/5 * * * * ?", command.PayCompare)
	c.AddFunc("* * 23 1 * ?", command.DeleteBillDir)

	c.Start()
	//加载会阻塞
	select {}
}
