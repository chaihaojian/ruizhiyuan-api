package models

import "time"

type Article struct {
	ID         int64     `json:"id" db:"id"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Outline    string    `json:"outline" db:"outline" binding:"required"`
	Author     string    `json:"author" db:"author" binding:"required"`
	Source     string    `json:"source" db:"source" binding:"required"`
	Content    string    `json:"text" db:"content" binding:"required"`
	Partition  string    `json:"partition" db:"partition" binding:"required"`
	Cover      string    `json:"cover" db:"cover" binding:"required"`
	IsShow     bool      `json:"is_show" db:"is_show"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
	Status     int       `json:"status" db:"status"`
}
