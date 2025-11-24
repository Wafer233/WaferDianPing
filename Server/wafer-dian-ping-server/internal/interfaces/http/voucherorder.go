package http

import "github.com/gin-gonic/gin"

type VoucherOrderHandler struct {
}

func NewVoucherOrderHandler() *VoucherOrderHandler {
	return &VoucherOrderHandler{}
}

func (h *VoucherOrderHandler) SeckillVoucher(c *gin.Context) {}
