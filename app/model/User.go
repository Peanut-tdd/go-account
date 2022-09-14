package model

import "time"

type User struct {
	ID        int64     `gorm:"primary"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
