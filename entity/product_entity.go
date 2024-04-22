package entity

import "time"

type Product struct {
	Id              int
	ProductName     string
	Stock           int
	Price           int
	Description     string
	CategoryProduct string
	PictureUrl      string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
