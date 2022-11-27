package models

import "time"

type RoleAllowed string

const (
	admin  RoleAllowed = "admin"
	member RoleAllowed = "member"
)

type Users struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	FullName  string
	Email     string
	Password  string
	Role      RoleAllowed
	CreatedAt time.Time
	UpdatedAt time.Time
	// Role      string `gorm:"type:enum('member','admin');default:'member'"`
}
