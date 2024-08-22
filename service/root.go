package service

import (
	"log"
	"server-api/config"
	"server-api/repository"
)

type Service struct {
	cfg        *config.Config
	repository *repository.Repository
}

func NewService(cfg *config.Config, repository *repository.Repository) (*Service, error) {
	s := &Service{cfg: cfg, repository: repository}

	return s, nil
}

func (s *Service) AddStock(name string) error {
	if err := s.repository.Mongo.AddStock(name); err != nil {
		log.Println("AddStock error: ", err)
		return err
	} else {
		return nil
	}
}
