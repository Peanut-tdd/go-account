package model

import "time"

type Project struct {
	ID               uint               `gorm:"primaryKey"`
	Name             string             `gorm:"type:varchar(200);not null;default ''"`
	VirtualCoinName  string             `gorm:"type:varchar(100);not null;default ''"`
	StartPayEq       int                `gorm:"not null,default 8"`
	Fee              int                `gorm:"not null;default 12"`
	Status           uint8              `gorm:"not null;default 1"`
	CreatedAt        time.Time          `gorm:"type:datetime"`
	UpdatedAt        time.Time          `gorm:"type:datetime"`
	ProjectAppConfig []ProjectAppConfig `gorm:"foreignKey:ProjectId"`
}

