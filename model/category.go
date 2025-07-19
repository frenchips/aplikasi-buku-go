package model

import (
	"time"
)

type Category struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	ModifiedAt time.Time `json:"modifiedAt"`
	ModifiedBy string    `json:"modifiedBy"`
}
