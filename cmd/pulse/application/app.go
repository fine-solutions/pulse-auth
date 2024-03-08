package application

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"pulse-auth/cmd/pulse/config"
	"pulse-auth/internal/authentication"
	"pulse-auth/internal/service/user"
	"pulse-auth/internal/storage/postgres"
	"pulse-auth/internal/token"
	"syscall"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	Closer *Closer
}

func New(cfg *config.Config, logger *zap.Logger) *App {
	return &App{
		Config: cfg,
		Logger: logger,
		Closer: NewCloser(logger, cfg.Application.GracefulShutdownTimeout, syscall.SIGINT, syscall.SIGTERM),
	}
}

func (a *App) Run() error {
	ctx, cancelFunction := context.WithCancel(context.Background())
	a.Closer.Add(func() error {
		cancelFunction()
		return nil
	})

	envStruct, err := a.constructEnv(ctx)
	if err != nil {
		return fmt.Errorf("construct env: %w", err)
	}

	httpServer := a.newHTTPServer(envStruct)
	a.Closer.Add(httpServer.GracefulStop()...)

	a.Closer.Run(httpServer.Run()...)
	a.Closer.Wait()
	return nil
}

type env struct {
	userService           user.Service
	authenticationService authentication.Service
}

func (a *App) constructEnv(ctx context.Context) (*env, error) {
	db, err := postgres.NewStorage(a.Logger, a.Config.Storage)
	if err != nil {
		return nil, fmt.Errorf("new storage: %w", err)
	}

	tokenGenerator := token.NewGenerator(a.Config.Application)
	userService := &user.ServiceImpl{
		Storage:        db,
		TokenGenerator: tokenGenerator,
		Logger:         a.Logger,
	}

	return &env{
		userService:           userService,
		authenticationService: authentication.Service{UserService: userService},
	}, nil
}
