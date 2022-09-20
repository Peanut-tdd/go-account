package model

import "time"

type Bill struct {
	//ID         uint32 `gorm:"primary_key"`
	Number     string `gorm:"type:varchar(50);index"`
	TradeNo    string `gorm:"type:varchar(50);index"`
	TradeAt    time.Time
	Amount     int
	PlatFormId int `gorm:"type:tinyint"`
}
