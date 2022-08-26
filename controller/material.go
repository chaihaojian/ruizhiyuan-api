package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ruizhiyuan/logic"
	"ruizhiyuan/models"
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
