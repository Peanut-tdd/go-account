package wechat

import (
	"account_check/app/utils"
	"fmt"
)

//Sign 微信签名加密
func Sign(request map[string]string) string {
	sign := utils.MD5Params(request, "636x44f9y3OqN65DRbrh4Zqydobt6MBW", nil)
	fmt.Print("签名加密：" + sign)

	return sign
}
