package network

import (
	"github.com/gin-gonic/gin"
	"log"
)

type API_REQUEST uint8

const (
	GET API_REQUEST = iota
	POST
	PUT
	DELETE
)

func (n *Network) register(path string, request API_REQUEST, h ...gin.HandlerFunc) gin.IRoutes {
	switch request {
	case GET:
		return n.engin.GET(path, h...)
	case POST:
		return n.engin.POST(path, h...)
	case PUT:
		return n.engin.PUT(path, h...)
	case DELETE:
		return n.engin.DELETE(path, h...)
	default:
		log.Println("API Register Failed", path)
		return nil
	}
}
