package models

import "time"

type File struct {
	FileID     int64     `json:"id" db:"id"`
	FileSize   int64     `json:"file_size" db:"file_size"`
	Status     int       `json:"status" db:"status"`
	Partition  string    `json:"partition" db:"partition" binding:"required"`
	FileSha1   string    `json:"file_sha1" binding:"required" db:"file_sha1"`
	FileName   string    `json:"file_name" binding:"required" db:"file_name"`
	FileAddr   string    `json:"file_addr" db:"file_addr"`
	UpDateTime time.Time `json:"update_time" db:"update_time"`
}
