package http

import "github.com/gin-gonic/gin"

type ShopHandler struct {
}

func NewShopHandler() *ShopHandler {
	return &ShopHandler{}
}

func (h *ShopHandler) QueryShopById(c *gin.Context) {}

func (h *ShopHandler) SaveShop(c *gin.Context) {}

func (h *ShopHandler) UpdateShop(c *gin.Context) {}

func (h *ShopHandler) QueryShopByType(c *gin.Context) {}

func (h *ShopHandler) QueryShopByName(c *gin.Context) {}
