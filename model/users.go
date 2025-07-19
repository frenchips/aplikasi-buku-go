package model

import "time"

type Users struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"createdAt"`
	CreateBy string    `json:"createdBy"`
}
