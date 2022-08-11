package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ruizhiyuan/logic"
	"ruizhiyuan/models"
)

func LoginHandler(c *gin.Context) {
	//获取参数及校验
	p := new(models.ParamLogin)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, "login with invalid param")
		return
	}

	//签发token
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "login failed")
		return
	}

	//返回响应
	ResponseSuccess(c, token)
	return
}
