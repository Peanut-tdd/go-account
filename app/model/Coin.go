package model

import "time"

type UserAccount struct {
	UserId             int
	Coin               int
	TotalAmount        int
	VipExpireTime      time.Time
	HalfPastExpireTime time.Time
	LastRechargeTime   time.Time
	CreatedAt          time.Time `gorm:"type:datetime"`
	UpdatedAt          time.Time `gorm:"type:datetime"`
}

type CoinFlow struct {
	ID         uint      `gorm:"primary_key"`
	TradeNo    string    `gorm:"type:varchar(50);index:index_trade_no;not null;default ''"`
	TradeAt    time.Time `gorm:"type:datetime"`
	UserId     int
	ProjectId  int
	PlatformId uint8 `gorm:"type:tinyint;not null;default 1"`
	Coin       int
	RealCoin   int
	Type       uint8
	Source     string
	SourceId   int
	Status     int
	CreatedAt  time.Time `gorm:"type:datetime"`
	UpdatedAt  time.Time `gorm:"type:datetime"`
}
