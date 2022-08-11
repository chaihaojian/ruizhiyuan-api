package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"ruizhiyuan/models"
)

func QueryAdmin(admin *models.Admin) error {
	sqlStr := "select id, name, password, status, last_active from admin where name = ? "
	err := db.Get(admin, sqlStr, admin.Name)
	if err == sql.ErrNoRows {
		zap.L().Error("admin not exist", zap.Error(err))
		return err
	}
	if err != nil {
		zap.L().Error("query admin failed", zap.Error(err))
		return err
	}
	return nil
}
