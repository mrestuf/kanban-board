package views

import "time"

type Register struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Login struct {
	Token string `json:"token"`
}

type UpdateUser struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

