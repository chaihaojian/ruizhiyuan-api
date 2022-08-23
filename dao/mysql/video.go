package mysql

import (
	"go.uber.org/zap"
	"ruizhiyuan/models"
)

func GetAllVideo() ([]models.Video, error) {
	sqlStr := "select `id`, `partition`, `interviewee`, `title`, `link`, `update_time`, `status` from video where status != 1"
	rows, err := db.Query(sqlStr)
	if err != nil {
		zap.L().Error("mysql get all video failed", zap.Error(err))
		return nil, err
	}

	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var v models.Video
		err := rows.Scan(&v.ID, &v.Partition, &v.Interviewee, &v.Title, &v.Link, &v.UpdateTime, &v.Status)
		if err != nil {
			zap.L().Error("mysql get all video failed", zap.Error(err))
			return nil, err
		}
		videos = append(videos, v)
	}
	return videos, nil
}

func AddVideo(v *models.Video) error {
	sqlStr := "insert into video(`id`, `partition`, `title`, `link`) values(?,?,?,?)"
	_, err := db.Exec(sqlStr, v.ID, v.Partition, v.Title, v.Link)
	if err != nil {
		zap.L().Error("mysql add video failed", zap.Error(err))
		return err
	}
	return nil
}

func DeleteVideo(vid int64) error {
	sqlStr := "delete from video where id = ?"
	_, err := db.Exec(sqlStr, vid)
	if err != nil {
		zap.L().Error("mysql delete video failed", zap.Error(err))
		return err
	}
	return nil
}

func SetVideoStatus(vid int64, status int) error {
	sqlStr := "update video set status = ? where id = ?"
	_, err := db.Exec(sqlStr, status, vid)
	if err != nil {
		zap.L().Error("mysql set video status failed", zap.Error(err))
		return err
	}
	return nil
}
