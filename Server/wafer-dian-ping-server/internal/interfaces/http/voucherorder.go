package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/application"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type VoucherOrderHandler struct {
	svc *application.VoucherOrderService
}

func NewVoucherOrderHandler(svc *application.VoucherOrderService) *VoucherOrderHandler {
	return &VoucherOrderHandler{
		svc: svc,
	}
}

func (h *VoucherOrderHandler) SeckillVoucher(c *gin.Context) {

	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, result.Fail("获取id失败"))
		return
	}

	id, _ := strconv.ParseInt(idStr, 10, 64)

	userId, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Fail("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()
	orderId, err := h.svc.SeckillService(ctx, id, userId.(int64))

	if err != nil || orderId == 0 {
		c.JSON(http.StatusInternalServerError, result.Fail("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.OkData(orderId))

}
