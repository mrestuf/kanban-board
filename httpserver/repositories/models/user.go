package models

import "time"

type User struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	FullName  string
	Email     string
	Password  string
	Role      string `gorm:"type:enum('member','admin');default:'member'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
