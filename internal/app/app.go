package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2pizzzza/IskenderBackend/api/user"
	"github.com/2pizzzza/IskenderBackend/internal/config"
	userHandler "github.com/2pizzzza/IskenderBackend/internal/http/user"
	userService "github.com/2pizzzza/IskenderBackend/internal/service/user"
	intPostgres "github.com/2pizzzza/IskenderBackend/internal/storage/postgres"
	"github.com/2pizzzza/IskenderBackend/pkg/httpserver"
	"github.com/2pizzzza/IskenderBackend/pkg/logger"
	pkgPostgres "github.com/2pizzzza/IskenderBackend/pkg/postgres"
	"go.uber.org/zap"
)

func New(cfg *config.Config) {

	ctx := context.Background()
	log, err := logger.New(cfg)
	if err != nil {
		fmt.Println("%w", err)
	}

	pg, err := pkgPostgres.New(ctx, cfg)
	if err != nil {
		log.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}

	_ = intPostgres.New(pg.Pool)

	userRepo := intPostgres.NewUserRepository(pg.Pool)
	userSrv := userService.NewUserService(userRepo)
	userH := userHandler.NewUserHandler(userSrv)

	defer pg.Close()

	application := httpserver.New(log, cfg.App.Host, cfg.App.Port)

	api.RegisterHandlers(application.App, userH)
	go func() {
		if err := application.Run(); err != nil {
			log.Error("Failed to run HTTP server", zap.Error(err))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("Stopping application", zap.String("signal", sign.String()))

	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	application.Stop(shutdownCtx)

}
