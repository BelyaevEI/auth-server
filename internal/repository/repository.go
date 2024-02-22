package repository

import (
	"github.com/BelyaevEI/auth-server/internal/config"
	"github.com/BelyaevEI/auth-server/internal/storage"
)

type Repositorer interface {
}

type repository struct {
	mongoDB storage.Storager
}

func NewRepository(cfg config.Config) (Repositorer, error) {

	mongoDB, err := storage.NewConnect(cfg)
	if err != nil {
		return nil, err
	}

	return &repository{mongoDB: mongoDB}, nil
}
