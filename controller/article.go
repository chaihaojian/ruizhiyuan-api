package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/url"
	"ruizhiyuan/logic"
	"ruizhiyuan/models"
	"strconv"
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

	if len(articles) == 0 {
		ResponseSuccess(c, 0)
		return
	}

	ResponseSuccess(c, articles)
	return
}

func DeleteArticleHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		zap.L().Error("invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err = logic.DeleteArticle(id); err != nil {
		zap.L().Error("delete article failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "delete article failed")
		return
	}

	ResponseSuccess(c, nil)
}

func UpdateArticleHandler(c *gin.Context) {
	article := new(models.Article)
	article.ID, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
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

	if err = logic.UpdateArticle(article, cover, header); err != nil {
		zap.L().Error("UpdateArticle failed", zap.Error(err))
		ResponseError(c, CodeError)
	}

	ResponseSuccess(c, nil)
}

func GetArticleCoverHandler(c *gin.Context) {
	path := c.Query("path")

	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", "attachment; filename="+url.QueryEscape(path))
	c.Writer.Header().Add("Content-Transfer-Encoding", "binary")
	c.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	c.Writer.Header().Add("response-type", "blob")

	c.File(viper.GetString("file.path") + path)
}
