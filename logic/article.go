package logic

import (
	"go.uber.org/zap"
	"mime/multipart"
	"ruizhiyuan/dao/mysql"
	"ruizhiyuan/models"
	"ruizhiyuan/pkg/snowflake"
	"ruizhiyuan/utils"
)

func AddArticleHandler(article *models.Article, cover multipart.File, header *multipart.FileHeader) error {
	//保存文章封面图片
	_, err := utils.FileUpLoad(cover, header)
	if err != nil {
		zap.L().Error("utils.FileUpLoad failed", zap.Error(err))
		return err
	}
	article.Cover = header.Filename
	article.ID = snowflake.GenID()

	if err = mysql.AddArticle(article); err != nil {
		zap.L().Error("mysql.AddArticle(article) failed", zap.Error(err))
		return err
	}

	return nil
}

func GetAllArticle() ([]models.Article, error) {
	articles, err := mysql.GetAllArticle()
	return articles, err
}

func DeleteArticle(id int64) error {
	return mysql.DeleteArticle(id)
}

func UpdateArticle(article *models.Article, cover multipart.File, header *multipart.FileHeader) error {
	//保存文章封面图片
	path, err := utils.FileUpLoad(cover, header)
	if err != nil {
		zap.L().Error("utils.FileUpLoad failed", zap.Error(err))
		return err
	}
	article.Cover = path

	if err = mysql.UpdateArticle(article); err != nil {
		zap.L().Error("mysql.UpdateArticle(article) failed", zap.Error(err))
		return err
	}

	return nil
}
