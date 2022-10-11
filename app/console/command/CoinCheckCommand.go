package command

import (
	"account_check/app/model"
	"account_check/bootstrap/driver"
	"fmt"
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

func Check() (isOk bool) {
	type UserCoin struct {
		UserId    uint
		TotalCoin uint
	}
	type UserAccountPart struct {
		UserId    uint
		TotalCoin uint `gorm:"column:coin"`
	}

	var groupUserCoin []UserAccountPart
	var incomeUserCoin []UserCoin
	var costUserCoin []UserCoin

	//账号剩余
	driver.GVA_DB.Model(&model.UserAccount{}).Order("user_id asc").Find(&groupUserCoin)
	//账号收入
	driver.GVA_DB.Model(&model.CoinFlow{}).Select("user_id,sum(real_coin) as total_coin").Where("type", 2).Group("user_id").Order("user_id asc").Find(&incomeUserCoin)
	//账号支出
	driver.GVA_DB.Model(&model.CoinFlow{}).Select("user_id,sum(real_coin) as total_coin").Where("type", 1).Group("user_id").Order("user_id asc").Find(&costUserCoin)
	if len(groupUserCoin) == 0 || len(incomeUserCoin) == 0 {
		return
	}

	var mapGroupUserCoin = make(map[uint]uint, 0)
	var mapIncomeUserCoin = make(map[uint]uint, 0)
	var mapCostUserCoin = make(map[uint]uint, 0)

	for _, value := range groupUserCoin {
		mapGroupUserCoin[value.UserId] = value.TotalCoin
	}

	for _, value := range incomeUserCoin {
		mapIncomeUserCoin[value.UserId] = value.TotalCoin
	}

	for _, value := range costUserCoin {
		mapCostUserCoin[value.UserId] = value.TotalCoin
	}

	var falseUserIds = make([]uint, 0)

	//逻辑判断
	for k, value := range mapIncomeUserCoin {
		if v, ok := mapCostUserCoin[k]; ok { //有消费
			if value != v+mapGroupUserCoin[k] {
				falseUserIds = append(falseUserIds, k)
			}

		} else { //无消费
			if value != mapGroupUserCoin[k] {
				falseUserIds = append(falseUserIds, k)
			}
		}
	}

	if len(falseUserIds) == 0 {
		isOk = true
	} else {
		isOk = false
		fmt.Println(falseUserIds)
	}

	return

}
