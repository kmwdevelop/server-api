package repository

import "server-api/config"

type Repository struct {
	cfg *config.Config
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	s := &Repository{cfg: cfg}

	return s, nil
}
