package models

import (
	"time"
)

type Categories struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	TaskId    int
	Tasks     *Task `gorm:"foreignKey:Id"`
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	// UserId    int
	// User      Users `gorm:"foreignKey:UserId"`
	// Task      views.TaskGetCategories
	// Task *Task
}
