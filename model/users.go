package model

import "time"

type User struct {
	ID        int       `json:"userId"`
	Name      string    `json:"name" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
