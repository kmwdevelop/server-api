package app

import (
	"server-api/config"
	"server-api/network"
	"server-api/repository"
	"server-api/service"
)

type App struct {
	cfg *config.Config

	network    *network.Network
	service    *service.Service
	repository *repository.Repository
}

func NewApp(cfg *config.Config) {
	a := &App{cfg: cfg}

	var err error

	if a.repository, err = repository.NewRepository(cfg); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service); err != nil {
		panic(err)
	} else {
		// TODO 서버 시작
	}
}
