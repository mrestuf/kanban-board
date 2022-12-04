package models

import "time"

type Task struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	Status      bool
	UserId      int
	User        Users `gorm:"foreignKey:UserId"`
	CategoryId  int
	Category    Categories `gorm:"foreignKey:CategoryId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
