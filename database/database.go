package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-user-service/config"
	"go.uber.org/zap"
)

func OpenConnect(cfg *config.DatabaseConfig) *sqlx.DB {
	db, err := sqlx.Connect(cfg.DbType, cfg.ConnectionString)
	if err != nil {
		logger.Error("Connect to database fail: ", zap.Error(err))
		return nil
	}
	return db
}
