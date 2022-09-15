package kuaishou

import (
	"account_check/app/utils"
	"log"
)

//GetBills 获得快手账单
func GetBills(request map[string]string) {
	//getToken(request)
	//return

	sign := utils.MD5Params(request, "DpgS_kpK93Nq5cJUsRMp2A", nil, "KS")
	request["sign"] = sign
	//获得url参数字段

	var queryForm = make(map[string]string)
	queryForm["app_id"] = request["app_id"]
	queryForm["access_token"] = "ChFvYXV0aC5hY2Nlc3NUb2tlbhIwgw5lycA7JOMGAJ2VU7H3tqlH9r0h4NZajlCKzvXgDjaofFkjoIRdpUyGkJBHogYQGhK2-ObXGzJKd7El7e1ZzY3H690iIKbeu3gcvM6etFlTTsV-jWsFDgecGy1b-6fZ1OrUfokgKAUwAQ" //getToken(request)
	//request["access_token"] = "ChFvYXV0aC5hY2Nlc3NUb2tlbhIwjQRCzKzDFUZcvobputGJDxLJHNY1VrjKW84lTh_yfUyp36TN1CU2yqNI53P0qUm0GhKLDWZIFONJcLwT16HoOA7b3moiINfxOPxEJFP1llu0PE-9VAjfEcBAFMZeE_wVcTPPYnnZKAUwAQ"

	res, _ := utils.HttpSendBodyResDownLoad("https://open.kuaishou.com/openapi/mp/developer/epay/query_bill", "post", request, queryForm, nil, "./test.zip", "")

	log.Print(res)
}

func getToken(request map[string]string) {
	//获取redis中的token，获取不到请求接口
	request["app_id"] = "ks695806146341101215"
	request["app_secret"] = "DpgS_kpK93Nq5cJUsRMp2A"
	request["grant_type"] = "client_credentials"
	res, _ := utils.HttpSendFormResJson("https://open.kuaishou.com/oauth2/access_token", "post", request, nil, "")
	log.Print(res.String())
}
