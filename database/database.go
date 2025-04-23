package database

import (
	"github.com/hmlylab/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	logger := logger.NewLogger()
	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logger.Error("failed to connect to database")
	}
	logger.Info("database connection opened")
	return DB
}
