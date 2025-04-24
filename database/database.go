package database

import (
	"github.com/hmlylab/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DbOptions struct {
	DSN          string
	PreparedStmt bool
}

func ConnectDB(options DbOptions) *gorm.DB {
	logger := logger.NewLogger()
	DB, err := gorm.Open(postgres.Open(options.DSN), &gorm.Config{
		Logger:      gormLogger.Default.LogMode(gormLogger.Info),
		PrepareStmt: options.PreparedStmt,
	})
	if err != nil {
		logger.Error("failed to connect to database")
	}
	logger.Info("database connection opened")
	return DB
}
