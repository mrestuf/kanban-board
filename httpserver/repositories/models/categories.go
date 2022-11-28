package models

import "time"

type Categories struct {
	Id        int
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
