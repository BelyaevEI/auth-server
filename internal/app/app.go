package app

import (
	"net/http"

	"github.com/BelyaevEI/auth-server/internal/initialization"
	"github.com/BelyaevEI/auth-server/internal/logger"
)

func NewApp() (*http.Server, error) {

	// Create new connect to logger
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	// Initialization additional entites
	init, err := initialization.New(log)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    init.Host + ":" + init.Port,
		Handler: init.Route,
	}
	return server, nil
}
