package rest

import "rest-api-template/internal/domain/repositories"

type RESTServer struct {
	usecases repositories.ContactUseCases
	logger   repositories.Logger
}

func NewRESTServer(
	usecases repositories.ContactUseCases,
	logger repositories.Logger,
) *RESTServer {
	return &RESTServer{
		usecases: usecases,
		logger:   logger,
	}
}

func (r *RESTServer) Start() error {
	r.logger.Info("Hi!")
	return nil
}

func (r *RESTServer) Stop() {}
