package initialization

import (
	"github.com/BelyaevEI/auth-server/internal/config"
	"github.com/BelyaevEI/auth-server/internal/handlers"
	"github.com/BelyaevEI/auth-server/internal/logger"
	"github.com/BelyaevEI/auth-server/internal/route"
	"github.com/go-chi/chi"
)

type Init struct {
	Host  string
	Port  string
	Route *chi.Mux
}

// Initialization all entites
func New(log *logger.Logger) (Init, error) {

	// Reading config file
	cfg, err := config.LoadConfig("../")
	if err != nil {
		log.Log.Error("read config file is fail: ", err)
		return Init{}, err
	}

	// Create new handlers
	handlers, err := handlers.New(log, cfg)
	if err != nil {
		log.Log.Error("initialization is failed ", err)
		return Init{}, err
	}

	// Create new router
	route := route.NewRoute(handlers)

	return Init{
		Route: route,
		Host:  cfg.Host,
		Port:  cfg.Port,
	}, nil
}
