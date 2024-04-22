package entity

import "time"

type User struct {
	Id        int
	Name      string
	UserName  string
	Email     string
	Password  string
	Gender    string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
