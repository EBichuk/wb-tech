package app

import (
	"calendar-service/internal/config"
	"calendar-service/internal/handler"
	"calendar-service/internal/handler/event"
	"calendar-service/internal/repository"
	"calendar-service/internal/server"
	"calendar-service/internal/service"
	"calendar-service/pkg/logger"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	logger     *slog.Logger
	httpServer *server.Server
}

func New() *App {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to init configs %v", err)
	}

	logger := logger.NewLogger()
	logger.Info("init logger success")

	repo := repository.New()
	repo.Full()
	service := service.New(repo)

	handlersEvent := event.New(service)
	router := handler.RegisterRouter(handlersEvent, logger)

	server := server.NewServer(router, cnf.Addr, cnf.ReadTimeout, cnf.WriteTimeout, cnf.IdleTimeout)
	return &App{
		logger:     logger,
		httpServer: server,
	}
}

func (a *App) Run() {
	wg := &sync.WaitGroup{}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Go(func() {
		if err := a.httpServer.StartHTTPServer(); err != nil {
			a.logger.Info("Server is stopped")
			return
		}
	})

	<-sigChan

	a.Stop()

	wg.Wait()
}

func (a *App) Stop() {
	if err := a.httpServer.Stop(); err != nil {
		a.logger.Error("Server stopped with error", "error", err.Error())
	}
}
