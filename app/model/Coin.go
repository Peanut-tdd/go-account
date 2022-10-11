package model

import "time"

type UserAccount struct {
	UserId             uint
	Coin               uint
	TotalAmount        uint
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
	UserId     uint      `gorm:"index:idx_user_id"`
	ProjectId  uint
	PlatformId uint8 	 `gorm:"type:tinyint;not null;default 1"`
	Coin       uint
	RealCoin   uint
	Type       uint8
	Source     string
	SourceId   uint
	Status     uint8
	CreatedAt  time.Time `gorm:"type:datetime"`
	UpdatedAt  time.Time `gorm:"type:datetime"`
}
