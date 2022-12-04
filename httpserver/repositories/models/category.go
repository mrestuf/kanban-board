package models

import (
	"time"

	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
)

type Categories struct {
	Id         int `gorm:"primaryKey;autoIncrement"`
	UserId     int `gorm:"foreignKey:Id"`
	CategoryId int `gorm:"foreignKey:Id"`
	TaskId     int `gorm:"foreignKey:Id"`
	Type       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Task       views.TaskGetCategories
}
