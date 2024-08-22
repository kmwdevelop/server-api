package types

type AddStockReq struct {
	Name string `json:"name" binding:"required"`
}

type LikeStockReq struct {
	Name string `json:"name" binding:"required"`
}

type DisLikeStockReq struct {
	Name string `json:"name" binding:"required"`
}
