package models

import (	
	"time"
)

type Categories struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

