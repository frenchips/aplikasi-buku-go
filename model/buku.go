package model

import "time"

type Buku struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Category    Category  `json:"category"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	ReleaseYear int       `json:"releaseYear"`
	Price       int       `json:"price"`
	TotalPage   int       `json:"totalPage"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   string    `json:"createdBy"`
	ModifiedAt  time.Time `json:"modifiedAt"`
	ModifiedBy  string    `json:"modifiedBy"`
}
