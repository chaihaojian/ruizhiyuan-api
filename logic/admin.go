package logic

import (
	"errors"
	"go.uber.org/zap"
	"ruizhiyuan/dao/mysql"
	"ruizhiyuan/models"
	"ruizhiyuan/pkg/jwt"
)

func Login(p *models.ParamLogin) (token string, err error) {
	//查询用户是否存在
	admin := &models.Admin{
		Name: p.Name,
	}
	err = mysql.QueryAdmin(admin)
	if err != nil {
		zap.L().Error("admin not exist", zap.Error(err))
		return "", err
	}

	//校验密码
	if p.Password != admin.Password {
		err = errors.New("wrong password")
		zap.L().Error("wrong password", zap.Error(err))
		return "", err
	}

	//签发Token
	token, err = jwt.GenToken(admin.ID, admin.Name)
	if err != nil {
		zap.L().Error("jwt.GenToken failed", zap.Error(err))
		return "", err
	}
	return token, nil
}
