package mysql

import (
	"go.uber.org/zap"
	"ruizhiyuan/models"
)

func AddMaterial(file *models.File) error {
	sqlStr := "insert into file (`id`, `file_sha1`, `file_name`, `file_size`, `partition`, `file_addr`) values (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, file.FileID, file.FileSha1, file.FileName, file.FileSize, file.Partition, file.FileAddr)
	if err != nil {
		zap.L().Error("mysql add material failed", zap.Error(err))
		return err
	}
	return nil
}

func GetAllMaterial() ([]models.File, error) {
	sqlStr := "select `id`, `partition`, `file_name`, `file_size`, `update_time` from file where status != 1"
	rows, err := db.Query(sqlStr)
	if err != nil {
		zap.L().Error("mysql get all material failed", zap.Error(err))
		return nil, err
	}

	defer rows.Close()

	var list []models.File
	for rows.Next() {
		var v models.File
		err = rows.Scan(&v.FileID, &v.Partition, &v.FileName, &v.FileSize, &v.UpDateTime)
		if err != nil {
			zap.L().Error("mysql get all material failed", zap.Error(err))
			return nil, err
		}
		list = append(list, v)
	}
	return list, nil
}
