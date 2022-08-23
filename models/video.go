package models

import "time"

type Video struct {
	ID          int64     `json:"id" db:"id"`
	Partition   string    `json:"partition" db:"partition" binding:"required"`
	Interviewee string    `json:"interviewee" db:"interviewee"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Link        string    `json:"link" db:"link" binding:"required"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
	Status      int       `json:"status" db:"status"`
}
