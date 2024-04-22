package dto

import (
	"mime/multipart"
	"time"
)

type PostRequest struct {
	ProductName     string                `form:"product_name"`
	Stock           int                   `form:"stock"`
	Price           int                   `form:"price"`
	Description     string                `form:"description"`
	CategoryProduct string                `form:"category_product"`
	PictureUrl      *multipart.FileHeader `form:"picture_url"`
}

type PostResponse struct {
	ID              int       `json:"id"`
	ProductName     string    `json:"product_name"`
	Stock           int       `json:"stock"`
	Price           int       `json:"price"`
	Description     string    `json:"description"`
	CategoryProduct string    `json:"category_product"`
	PictureUrl      string    `json:"picture_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
