package kuaishou

import (
	"account_check/app/utils"
	"account_check/app/vo"
	"account_check/bootstrap/driver"
	"path"
	"time"
)

const KS_TOKEN_KEY = "ks:token"

//GetBills 获得快手账单
func GetBills(request map[string]string) {
	sign := utils.MD5Params(request, driver.GVA_VP.GetString("ks.app_secret"), nil, "KS")
	request["sign"] = sign
	//获得url参数字段
	var queryForm = make(map[string]string)
	queryForm["app_id"] = request["app_id"]
	queryForm["access_token"] = getToken(request)
	//获得下载账单
	filepath := utils.CsvFilePath(2)
	//生成下载文件
	utils.HttpSendBodyResDownLoad("https://open.kuaishou.com/openapi/mp/developer/epay/query_bill", "post", request, queryForm, nil, filepath, "")
	//解压账单
	utils.Unzip(filepath, utils.CsvFileDir(2))
	//获取文件夹下所有的文件
	files, _ := utils.TPFuncReadDirFiles(utils.CsvFileDir(2))
	for _, file := range files {
		//获得当前的csv文件
		if path.Ext(file) == ".csv" {
			//todo 读取csv文件并插入数据库中
			//log.Println(file)
		}
	}

}

//getToken 获得快手token
func getToken(request map[string]string) string {
	token := utils.RedisGet(KS_TOKEN_KEY)
	if token == nil {
		request["app_id"] = driver.GVA_VP.GetString("ks.app_id")
		request["app_secret"] = driver.GVA_VP.GetString("ks.app_secret")
		request["grant_type"] = "client_credentials"
		res, _ := utils.HttpSendFormResJson("https://open.kuaishou.com/oauth2/access_token", "post", request, nil, vo.GetAccessToken{})
		result := res.Result().(*vo.GetAccessToken)
		utils.RedisSet(KS_TOKEN_KEY, result.AccessToken, time.Duration(result.ExpiresIn)*time.Second)

		return result.AccessToken
	}

	return utils.GetInterfaceToString(token)
}
