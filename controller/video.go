package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ruizhiyuan/logic"
	"ruizhiyuan/models"
	"strconv"
)

func GetAllVideoHandler(c *gin.Context) {
	videos, err := logic.GetAllVideo()
	if err != nil {
		zap.L().Error("get all video failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "get all video failed")
		return
	}

	ResponseSuccess(c, videos)
	return
}

func AddVideoHandler(c *gin.Context) {
	//获取参数及校验
	v := new(models.Video)
	err := c.ShouldBindJSON(&v)
	if err != nil {
		zap.L().Error("invalid video param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "invalid video param")
		return
	}

	if err = logic.AddVideo(v); err != nil {
		zap.L().Error("add video failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "add video failed")
		return
	}

	ResponseSuccess(c, nil)
	return
}

func DeleteVideoHandler(c *gin.Context) {
	//获取参数
	idStr := c.Query("id")
	vid, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid video id", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "invalid video id")
		return
	}
	//删除数据(伪删除)
	if err := logic.DeleteVideo(vid); err != nil {
		zap.L().Error("logic.DeleteVideo(vid) failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "delete video failed")
		return
	}
	ResponseSuccess(c, nil)
	return
}
