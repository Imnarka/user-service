package app

import (
	"context"
	"github.com/Imnarka/user-service/internal/config"
	"github.com/Imnarka/user-service/internal/logger"
	"github.com/Imnarka/user-service/internal/transport/grpc"
	"gorm.io/gorm"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	server *grpc.Server
	db     *gorm.DB
	logger *logger.Logger
	cfg    *config.Config
}

func NewApp(server *grpc.Server, db *gorm.DB, logger *logger.Logger, cfg *config.Config) *App {
	return &App{
		server: server,
		db:     db,
		logger: logger,
		cfg:    cfg,
	}
}

func (app *App) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", ":"+app.cfg.GRPCPort)
	if err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	serverErrChan := make(chan error, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		app.logger.Info("Старт gRPC сервера...")
		if err := app.server.GrpcServer.Serve(listener); err != nil {
			serverErrChan <- err
		}
	}()

	select {
	case err := <-serverErrChan:
		app.logger.WithError(err).Error("gRPC сервер упал")
		return err
	case sig := <-sigChan:
		app.logger.WithField("signal", sig).Info("Принят сигнал завершения")
	case <-ctx.Done():
		app.logger.Info("Контекст отменен, инициализация shutdown")
	}

	// Выполняем graceful shutdown
	if err := app.Shutdown(context.Background(), &wg); err != nil {
		app.logger.WithError(err).Error("Shutdown не удался")
		return err
	}

	return nil
}

func (app *App) Shutdown(ctx context.Context, wg *sync.WaitGroup) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	app.logger.Info("Остановка gRPC сервера")
	app.server.GrpcServer.GracefulStop()
	app.logger.Info("gRPC сервер остановлен")

	wg.Wait()

	app.logger.Info("Закрытие соединения с БД")
	sqlDB, err := app.db.DB()
	if err != nil {
		app.logger.WithError(err).Error("Не удалось получить инстанс БД для закрытия соединения")
		return err
	}
	if err := sqlDB.Close(); err != nil {
		app.logger.WithError(err).Error("Не удалось закрыть соединение")
		return err
	}
	app.logger.Info("Соединение закрыто")

	// Проверяем, завершился ли shutdown вовремя
	select {
	case <-shutdownCtx.Done():
		app.logger.Warn("Shutdown timeout")
		return shutdownCtx.Err()
	default:
		app.logger.Info("Shutdown завершился успешно")
		return nil
	}
}
