package models

import "time"

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

type Users struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	FullName  string
	Email     string
	Password  string
	Role      Role `gorm:"type:role;default:'member'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
