package mysql

import (
	"go.uber.org/zap"
	"ruizhiyuan/models"
)

func QueryFile(file *models.File) error {
	sqlStr := "select `file_name`, `file_size`, `partition`, `file_addr`, `update_time` from file where id = ?"
	err := db.QueryRow(sqlStr, file.FileID).Scan(&file.FileName, &file.FileSize, &file.Partition, &file.FileAddr, &file.UpDateTime)
	if err != nil {
		zap.L().Error("db.Query failed", zap.Error(err))
		return err
	}
	return nil
}
