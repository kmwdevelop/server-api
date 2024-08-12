package network

import (
	"server-api/config"
	"server-api/service"
)

type Network struct {
	cfg     *config.Config
	service *service.Service
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	n := &Network{cfg: cfg, service: service}

	return n, nil
}
