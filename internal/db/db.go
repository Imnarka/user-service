package db

import (
	"fmt"
	"github.com/Imnarka/user-service/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

const (
	maxRetries = 10
)

func InitDB(cfg *DatabaseConfig, logger *logger.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	var db *gorm.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		logger.WithError(err).Debug("Не удалось подключиться к БД, попытка %d: %v", i+1)
		time.Sleep(2 * time.Second)
	}
	return db, err
}
