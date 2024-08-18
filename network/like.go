package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type like struct {
	network *Network
}

func newLike(network *Network) {
	l := &like{network: network}

	// TODO Like 요청에 대해 라우터 등록 필요

	basePath := "/like"

	network.register(basePath+"/add", POST, l.addStock)
	network.register(basePath+"/add", POST, l.likeStock)
	network.register(basePath+"/add", POST, l.disLikeStock)

	fmt.Println(l)
}

func (l *like) addStock(c *gin.Context) {

}

func (l *like) likeStock(c *gin.Context) {

}

func (l *like) disLikeStock(c *gin.Context) {

}
