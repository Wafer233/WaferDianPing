package http

import "github.com/gin-gonic/gin"

type VoucherHandler struct {
}

func NewVoucherHandler() *VoucherHandler {
	return &VoucherHandler{}
}

func (h *VoucherHandler) AddVoucher(c *gin.Context) {}

func (h *VoucherHandler) AddSeckillVoucher(c *gin.Context) {}

func (h *VoucherHandler) QueryVoucherOfShop(c *gin.Context) {}
