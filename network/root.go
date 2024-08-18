package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-api/config"
	"server-api/service"
)

type Network struct {
	cfg     *config.Config
	service *service.Service
	engin   *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	n := &Network{cfg: cfg, service: service, engin: gin.New()}

	// TODO Router
	n.register("/health-check", GET, healthCheck)
	newLike(n)
	return n, nil
}

func (n *Network) Run() error {
	return n.engin.Run(n.cfg.Network.Port)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "HealthCheck")
}
