package network

import (
	"fmt"
)

type like struct {
	network *Network
}

func newLike(network *Network) {
	l := &like{network: network}

	// TODO Like 요청에 대해 라우터 등록 필요

	fmt.Println(l)
}
