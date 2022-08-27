package logic

import (
	"go.uber.org/zap"
	"mime/multipart"
	"ruizhiyuan/dao/mysql"
	"ruizhiyuan/models"
	"ruizhiyuan/pkg/snowflake"
	"ruizhiyuan/utils"
)

func AddMaterial(file *models.File, f multipart.File, header *multipart.FileHeader) error {
	//保存文件
	addr, err := utils.FileUpLoad(f, header)
	if err != nil {
		zap.L().Error("utils.FileUpLoad failed", zap.Error(err))
		return err
	}

	//生成文件信息
	file.FileID = snowflake.GenID()
	file.FileAddr = addr
	file.FileName = header.Filename
	file.FileSha1 = utils.GetFileSha1(addr)
	file.FileSize = utils.GetFileSize(addr)

	//将文件信息保存进数据库
	if err = mysql.AddMaterial(file); err != nil {
		zap.L().Error("mysql.AddFile(file) failed", zap.Error(err))
		return err
	}
	return nil
}

func GetAllMaterial() ([]models.File, error) {
	return mysql.GetAllMaterial()
}

func DeleteMaterial(id int64) error {
	return mysql.DeleteMaterial(id)
}

func DownloadMaterial(material *models.File) error {
	return mysql.QueryFile(material)
}
