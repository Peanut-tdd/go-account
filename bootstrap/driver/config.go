package driver

import (
	"account_check/config"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Config *viper.Viper

var AllConfig *config.Config

//配置文件加载
func InitConfig() {
	//初始化viper
	Config := viper.New()
	//设置文件名
	Config.SetConfigName("application" + getProfiles())
	//设置文件类型
	Config.SetConfigType("yaml")
	//设置文件所在的目录
	Config.AddConfigPath("./")
	if err := Config.ReadInConfig(); err != nil {
		fmt.Println("init fail:", err.Error())
	}
	GVA_VP = Config

	err := Config.Unmarshal(&AllConfig)
	if err != nil {
		log.Print(err.Error())
	}

	mysql := AllConfig.GetKeyValue("Mysql", nil)

	fmt.Println(mysql)

}

//获得当前环境对应的配置信息
func getProfiles() string {
	var profiles string
	flag.StringVar(&profiles, "profiles", "", "配置文件信息，默认为空")
	flag.Parse()
	if profiles == "" {
		return ""
	}

	return "-" + profiles
}
