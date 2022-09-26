package kuaishou

import (
	"account_check/app/model"
	"account_check/app/utils"
	"account_check/app/vo"
	"account_check/bootstrap/driver"
	"encoding/csv"
	"fmt"
	"os"
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
	filepath := utils.CsvFilePath(2, 1)
	//生成下载文件
	utils.HttpSendBodyResDownLoad("https://open.kuaishou.com/openapi/mp/developer/epay/query_bill", "post", request, queryForm, nil, filepath, "")
	//解压账单

	err := utils.Unzip(filepath, utils.CsvFileDir(2, 1))
	if err == nil {
		//todo 打印日志
		return
	}
	//获取文件夹下所有的文件
	files, _ := utils.TPFuncReadDirFiles(utils.CsvFileDir(2, 1))
	for _, file := range files {
		//获得当前的csv文件
		if path.Ext(file) == ".csv" {
			//todo 读取csv文件并插入数据库中
			readBillCsv(file)
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

func readBillCsv(filepath string) {
	file, err := os.Open(filepath) //读取文件
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	csvdata, err := reader.ReadAll() //读取全部数据
	//账单数据
	var bills []model.OrderBill
	var numbers []string
	var numbersModel []string

	for index, item := range csvdata { //按行打印数据
		if index < 1 {
			continue
		}
		//获得所有
		numbers = append(numbers, item[4])
		bills = append(bills, model.OrderBill{
			Number:     item[4],
			TradeNo:    item[5],
			TradeAt:    utils.StringToTime(item[1]),
			Amount:     utils.GetInterfaceToInt(item[6]),
			PlatformId: 2,
		})
	}
	if len(numbers) <= 0 {
		return
	}
	driver.GVA_DB.Model(&model.OrderBill{}).Where("number in ?", numbers).Pluck("number", &numbersModel)
	//比较两数组的不同
	newNumber, _ := utils.Arrcmp(numbers, numbersModel)
	if len(newNumber) <= 0 {
		return
	}
	var newBills []model.OrderBill
	for _, item := range bills {
		for _, value := range newNumber {
			if item.Number == value {
				newBills = append(newBills, item)
			}
		}
	}

	//将订单插入数据库中
	if len(newBills) > 0 {
		driver.GVA_DB.Create(newBills)
	}
}
