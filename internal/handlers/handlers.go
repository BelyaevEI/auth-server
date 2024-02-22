package handlers

import (
	"net/http"

	"github.com/BelyaevEI/auth-server/internal/config"
	"github.com/BelyaevEI/auth-server/internal/logger"
	"github.com/BelyaevEI/auth-server/internal/repository"
)

type Handler interface {
	CreateNewToken(writer http.ResponseWriter, request *http.Request)
}

type Handlers struct {
	log        *logger.Logger
	repository repository.Repositorer
}

func New(log *logger.Logger, cfg config.Config) (*Handlers, error) {

	// Create new repository
	repository, err := repository.NewRepository(cfg)
	if err != nil {
		return &Handlers{}, err
	}

	return &Handlers{
		log:        log,
		repository: repository,
	}, nil
}
