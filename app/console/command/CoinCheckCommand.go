package command

import (
	"account_check/app/model"
	"account_check/bootstrap/driver"
	"log"
)

func CoinCheck() {
	coin := getCoinCount()
	payCoin := getPayCoinCount()
	incomeCoin := getIncomeCoinCount()
	log.Println(coin, payCoin, incomeCoin)
	if coin != (incomeCoin - payCoin) {
		//发送钉钉消息
	}
}

func UserCoinCheck() {

}

func getCoinCount() int {
	var coin int
	driver.GVA_DB.Model(model.UserAccount{}).Pluck("sum(coin) as coin", &coin)

	return coin
}

func getPayCoinCount() int {
	var coin int
	driver.GVA_DB.Model(model.CoinFlow{}).Where("type = ?", 1).Pluck("sum(real_coin) as coin", &coin)

	return coin
}

func getIncomeCoinCount() int {
	var coin int
	driver.GVA_DB.Model(model.CoinFlow{}).Where("type = ?", 2).Pluck("sum(real_coin) as coin", &coin)

	return coin
}
