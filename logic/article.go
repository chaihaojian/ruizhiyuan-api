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
	path, err := utils.FileUpLoad(cover, header)
	if err != nil {
		zap.L().Error("utils.FileUpLoad failed", zap.Error(err))
		return err
	}
	article.Cover = path
	article.ID = snowflake.GenID()

	if err = mysql.AddArticle(article); err != nil {
		zap.L().Error("mysql.AddArticle(article) failed", zap.Error(err))
		return err
	}

	return nil
}
