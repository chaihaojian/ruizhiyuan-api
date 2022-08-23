package logic

import (
	"ruizhiyuan/dao/mysql"
	"ruizhiyuan/models"
	"ruizhiyuan/pkg/snowflake"
)

func GetAllVideo() ([]models.Video, error) {
	videos, err := mysql.GetAllVideo()
	return videos, err
}

func AddVideo(v *models.Video) error {
	//生成video id
	v.ID = snowflake.GenID()
	return mysql.AddVideo(v)
}

func DeleteVideo(vid int64) error {
	return mysql.SetVideoStatus(vid, 1)
}
