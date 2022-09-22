package model

import "time"

type UserTest struct {
	ID        uint     `gorm:"primary_key"`
	Username  string
	CreatedAt time.Time `gorm:"autoCreateTime;type:datetime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;type:datetime"`
}
