package mysql

import (
	"go.uber.org/zap"
	"ruizhiyuan/models"
)

func AddArticle(a *models.Article) error {
	sqlStr := "insert into article(`id`, `partition`, `title`, `outline`, `author`, `source`, `content`, `cover`) values(?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sqlStr, a.ID, a.Partition, a.Title, a.Outline, a.Author, a.Source, a.Content, a.Cover)
	if err != nil {
		zap.L().Error("AddArticle db.Exec failed", zap.Error(err))
		return err
	}
	return nil
}

func GetAllArticle() ([]models.Article, error) {
	sqlStr := "select *from article where status != 1"
	rows, err := db.Query(sqlStr)
	if err != nil {
		zap.L().Error("mysql get all article failed", zap.Error(err))
		return nil, err
	}

	defer rows.Close()

	var article []models.Article
	for rows.Next() {
		var v models.Article
		err := rows.Scan(&v.ID, &v.Partition, &v.Author, &v.Source, &v.Title, &v.Outline, &v.Content, &v.Cover, &v.IsShow, &v.UpdateTime, &v.Status)
		if err != nil {
			zap.L().Error("mysql get all article failed", zap.Error(err))
			return nil, err
		}
		article = append(article, v)
	}
	return article, nil
}

func DeleteArticle(id int64) error {
	sqlStr := "delete from article where id = ?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		zap.L().Error("DeleteArticle failed", zap.Error(err))
		return err
	}
	return nil
}

func UpdateArticle(a *models.Article) error {
	sqlStr := "update article set `partition`=?, `title`=?, `outline`=?, `author`=?, `source`=?, `content`=?, `cover`=? where id= ?"
	_, err := db.Exec(sqlStr, a.Partition, a.Title, a.Outline, a.Author, a.Source, a.Content, a.Cover, a.ID)
	if err != nil {
		zap.L().Error("UpdateArticle db.Exec failed", zap.Error(err))
		return err
	}
	return nil
}
