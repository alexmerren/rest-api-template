package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/domain/repositories"
	"time"
)

const (
	readTimeoutInSeconds   = 5
	writeTimeoutInSeconds  = 10
	serverTimeoutInSeconds = 5
)

type RESTServer struct {
	usecases   repositories.ContactUseCases
	logger     repositories.Logger
	httpServer *http.Server
}

func NewRESTServer(
	usecases repositories.ContactUseCases,
	logger repositories.Logger,
	port int,
) *RESTServer {
	return &RESTServer{
		usecases: usecases,
		logger:   logger,
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			ReadTimeout:  readTimeoutInSeconds * time.Second,
			WriteTimeout: writeTimeoutInSeconds * time.Second,
		},
	}
}

func (s *RESTServer) Start() error {
	s.logger.Info(fmt.Sprintf("starting HTTP server %s", s.httpServer.Addr))
	if err := s.mapRoutes(); err != nil {
		s.logger.Error(err)
		return entities.NewInternalError("could not map routes", err)
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error(err)
		}
	}()
	return nil
}

func (s *RESTServer) Stop() error {
	s.logger.Debug("stopping HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), serverTimeoutInSeconds*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Info("server shutdown failed")
		return err
	}
	return nil
}
