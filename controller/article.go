package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ruizhiyuan/logic"
	"ruizhiyuan/models"
)

func AddArticleHandler(c *gin.Context) {
	article := new(models.Article)
	article.Title = c.PostForm("title")
	article.Outline = c.PostForm("outline")
	article.Author = c.PostForm("author")
	article.Source = c.PostForm("source")
	article.Partition = c.PostForm("partition")
	article.Content = c.PostForm("text")
	cover, header, err := c.Request.FormFile("cover")
	if err != nil {
		zap.L().Error("c.Request.FormFile failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
	}

	if err := logic.AddArticleHandler(article, cover, header); err != nil {
		zap.L().Error("logic.FileUpLoad failed", zap.Error(err))
		ResponseError(c, CodeError)
	}

	ResponseSuccess(c, nil)
}

func GetAllArticleHandler(c *gin.Context) {
	articles, err := logic.GetAllArticle()
	if err != nil {
		zap.L().Error("get all video failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "get all video failed")
		return
	}

	ResponseSuccess(c, articles)
	return
}
