package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/url"
	"ruizhiyuan/logic"
	"ruizhiyuan/models"
	"strconv"
)

func AddMaterialHandler(c *gin.Context) {
	//获取参数
	f, header, err := c.Request.FormFile("file")
	if err != nil {
		zap.L().Error("invalid file param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "invalid file param")
		return
	}
	file := new(models.File)
	file.Partition = c.PostForm("partition")

	if err = logic.AddMaterial(file, f, header); err != nil {
		zap.L().Error("file upload failed", zap.Error(err))
		ResponseError(c, CodeError)
		return
	}

	ResponseSuccess(c, file)
}

func GetAllMaterialHandler(c *gin.Context) {
	list, err := logic.GetAllMaterial()
	if err != nil {
		zap.L().Error("logic.GetAllMaterial failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "get material list failed")
		return
	}

	ResponseSuccess(c, list)
}

func DeleteMaterialHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		zap.L().Error("invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err = logic.DeleteMaterial(id); err != nil {
		zap.L().Error("delete article failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "delete article failed")
		return
	}

	ResponseSuccess(c, nil)
}

func DownloadMaterialHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	material := new(models.File)
	material.FileID = id

	if err := logic.DownloadMaterial(material); err != nil {
		zap.L().Error("logic.DownloadMaterial failed", zap.Error(err))
		ResponseError(c, CodeError)
		return
	}

	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", "attachment; filename="+url.QueryEscape(material.FileName))
	c.Writer.Header().Add("Content-Transfer-Encoding", "binary")
	c.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	c.Writer.Header().Add("response-type", "blob")

	c.File(material.FileAddr)
}
