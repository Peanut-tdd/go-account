package command

import (
	"account_check/app/model"
	"account_check/bootstrap/driver"
	"fmt"
)

//获取支付配置
func GetPayConfig() []model.Project {
	//获取已支付的项目id
	var ProjectIds []int
	driver.GVA_DB.Model(&model.Orders{}).Distinct().Pluck("project_id", &ProjectIds)

	var PayConfigs []model.Project
	driver.GVA_DB.Model(&model.Project{}).Preload("ProjectAppConfig").Where("id in ?", ProjectIds).Order("id asc").Find(&PayConfigs)

	return PayConfigs

}

func GetConfigByQueryParmas(projectId string, platformId string, payChannel string) model.ProjectAppConfig {

	var payConfig model.ProjectAppConfig

	//var queryParams = make(map[string]interface{},0)
	queryParams := map[string]interface{}{
		"project_id":  projectId,
		"platform_id": platformId,
		"pay_channel": payChannel,
	}

	fmt.Println(queryParams)
	driver.GVA_DB.Model(&model.ProjectAppConfig{}).Where(queryParams).Find(&payConfig)

	return payConfig
}
