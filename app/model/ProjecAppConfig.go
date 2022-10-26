package model

import "time"

type ProjectAppConfig struct {
	ProjectId       uint      `gorm:"primaryKey;autoIncrement:false"`
	PlatformId      uint8     `gorm:"primaryKey;autoIncrement:false;not null;default 1"`
	AppSign         string    `gorm:"type:varchar(255);not null; default ''"`
	AppId           string    `gorm:"type:varchar(100);not null;default ''"`
	AppSecret       string    `gorm:"type:varchar(100); not null;default ''"`
	MchId           string    `gorm:"type:varchar(32); not null;default ''"`
	Key             string    `gorm:"type:varchar(100);not null;default ''"`
	Salt            string    `gorm:"type:varchar(100);not null;default ''"`
	Token           string    `gorm:"type:varchar(100);not null;default ''"`
	PrivateKey      string    `gorm:"type:text"`
	AlipayPublicKey string    `gorm:"type:text"`
	NotifyUrl       string    `gorm:"type:varchar(255) not null ;default''"`
	Referer         string    `gorm:"type:varchar(255);not null;default ''"`
	RedirectUrl     string    `gorm:"type:varchar(255);not null;default ''"`
	PayChannel      uint8     `gorm:"primaryKey;autoIncrement:false;not null;default 1"`
	Status          uint8     `gorm:"not null;default 1"`
	CreatedAt       time.Time `gorm:"type:datetime"`
	UpdatedAt       time.Time `gorm:"type:datetime"`
}
