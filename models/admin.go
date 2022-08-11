package models

import "time"

type ParamLogin struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Admin struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Password   string    `json:"password" db:"password"`
	Status     int       `json:"status" db:"status"`
	LastActive time.Time `json:"last_active" db:"last_active"`
}
