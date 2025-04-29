//go:build wireinject

package di

import (
	"github.com/Imnarka/user-service/internal/app"
	"github.com/Imnarka/user-service/internal/config"
	"github.com/Imnarka/user-service/internal/db"
	"github.com/Imnarka/user-service/internal/logger"
	"github.com/Imnarka/user-service/internal/transport/grpc"
	"github.com/Imnarka/user-service/internal/users"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type AppComponents struct {
	App    *app.App
	DB     *gorm.DB
	Logger *logger.Logger
	Config *config.Config
}

func provideDBConfig(cfg *config.Config) *db.DatabaseConfig {
	return &db.DatabaseConfig{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  cfg.DBSSLMode,
	}
}

func InitializeGRPCServer(cfg *config.Config) (*AppComponents, error) {
	wire.Build(
		loggerSet,
		databaseSet,
		userSet,
		transportSet,
		app.NewApp,
		wire.Struct(new(AppComponents), "*"),
	)
	return nil, nil
}

var loggerSet = wire.NewSet(
	logger.InitLogger,
)

var databaseSet = wire.NewSet(
	provideDBConfig,
	db.InitDB,
)

var userSet = wire.NewSet(
	users.NewUserRepository,
	users.NewService,
)

var transportSet = wire.NewSet(
	grpc.NewHandler,
	grpc.NewServer,
)
