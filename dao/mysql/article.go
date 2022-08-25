package mysql

import (
	"go.uber.org/zap"
	"ruizhiyuan/models"
)

func AddArticle(a *models.Article) error {
	sqlStr := "insert into article(`id`, `partition`, `title`, `outline`, `author`, `source`, `content`, `cover`) values(?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sqlStr, a.ID, a.Partition, a.Title, a.Outline, a.Author, a.Source, a.Text, a.Cover)
	if err != nil {
		zap.L().Error("AddArticle db.Exec failed", zap.Error(err))
		return err
	}
	return nil
}
