package model

import "time"

type OrderBill struct {
	ID         uint      `gorm:"primary_key"`
	ProjectId  uint      `gorm:"not null;default 1"`
	Number     string    `gorm:"type:varchar(50);index:idx_number;not null;default ''"`
	TradeNo    string    `gorm:"type:varchar(50);index:index_trade_no;not null;default ''"`
	TradeAt    time.Time `gorm:"type:datetime"`
	Amount     int       `gorm:"type:int;not null;default 0"`
	PlatformId uint8     `gorm:"type:tinyint;not null;default 1"`
	PayChannel uint8     `gorm:"type:tinyint;not null;default 1"`
	CreatedAt  time.Time `gorm:"type:datetime"`
	UpdatedAt  time.Time `gorm:"type:datetime"`
}

type Orders struct {
	ID             uint `gorm:"primary_key"`
	ProjectId      int
	Number         string `gorm:"type:varchar(50);index:idx_number;not null;default ''"`
	TradeNo        string `gorm:"type:varchar(50);index:index_trade_no;not null;default ''"`
	UserId         int
	PChargeId      int
	PChargeName    int
	ChargeId       int
	Money          int `gorm:"type:int;not null;default 0"`
	TotalAmount    int `gorm:"type:int;not null;default 0"`
	Days           int
	ActivityPrice  int
	GetCoin        int
	FreeCoin       int
	Status         int
	PaySuccessTime time.Time `gorm:"type:datetime"`
	PlatformId     uint8     `gorm:"type:tinyint;not null;default 1"`
	PayChannel     uint8     `gorm:"type:tinyint;not null;default 1"`
	CreatedAt      time.Time `gorm:"type:datetime"`
	UpdatedAt      time.Time `gorm:"type:datetime"`
}
