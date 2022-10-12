package command

import (
	"account_check/app/model"
	"account_check/bootstrap/driver"
	"fmt"
	"log"
	"math"
	"strconv"
)





func CoinCheckMessage() {
	var message string
	count := PageCheck()

	if count != 0 {
		message += "\n\n--------------------\n\n异常账户数："+ strconv.Itoa(count)
	}
	title := "账户虚拟币对账异常"

	if message != "" {
		sendMessage(title+message, title)
	}

}






func CoinCheck() {
	coin := getCoinCount()
	payCoin := getPayCoinCount()
	incomeCoin := getIncomeCoinCount()
	log.Println(coin, payCoin, incomeCoin)
	if coin != (incomeCoin - payCoin) {
		//发送钉钉消息
	}
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


/**
分页查询
*/
func PageCheck() int {

	var total int64
	const PAGESIZE = 10

	driver.GVA_DB.Model(&model.CoinFlow{}).Where("type", 2).Distinct("user_id").Count(&total)
	if total == 0 {
		return 0
	}

	var maxPage = int(math.Ceil(float64(total) / float64(PAGESIZE)))
	type UserCoin struct {
		UserId uint
		Coin   uint
		Income uint `gorm:"column:income_coin"`
		Out    uint `gorm:"column:out_coin"`
	}

	var result []UserCoin
	var errUserIds = make([]uint, 0)

	for i := 1; i <= maxPage; i++ {
		var offset = (i - 1) * PAGESIZE

		iquery := driver.GVA_DB.Model(&model.CoinFlow{}).Select("user_id,sum(real_coin) as real_coin").Where("type", 2).Group("user_id")
		oquery := driver.GVA_DB.Model(&model.CoinFlow{}).Select("user_id,sum(real_coin) as real_coin").Where("type", 1).Group("user_id")
		driver.GVA_DB.Model(&model.UserAccount{}).Select("user_account.user_id,user_account.coin,i.real_coin as income_coin,o.real_coin as out_coin").Joins("join (?) i on i.user_id=user_account.user_id", iquery).
			Joins("left join (?) o on o.user_id=user_account.user_id", oquery).Order("user_account.user_id asc").Limit(PAGESIZE).Offset(offset).Scan(&result)

		for _, value := range result {
			if value.Out == 0 { //无消费
				if value.Coin != value.Income {
					errUserIds = append(errUserIds, value.UserId)
				}
			} else { //有消费
				if value.Coin != value.Income-value.Out {
					errUserIds = append(errUserIds, value.UserId)
				}
			}
		}

	}
	return len(errUserIds)

}
