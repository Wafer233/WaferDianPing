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

type VoucherHandler struct {
	svc *application.VoucherService
}

func NewVoucherHandler(
	svc *application.VoucherService,
) *VoucherHandler {
	return &VoucherHandler{
		svc: svc,
	}
}

func (h *VoucherHandler) AddVoucher(c *gin.Context) {

	dto := application.VoucherDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail("输入有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	id, err := h.svc.CreateVoucher(ctx, &dto)
	if err != nil || id == 0 {
		c.JSON(http.StatusInternalServerError, result.Fail("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.OkData(id))

}

func (h *VoucherHandler) AddSeckillVoucher(c *gin.Context) {

	dto := application.VoucherDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail("输入有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	id, err := h.svc.CreateSeckillVoucher(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.OkData(id))
}

func (h *VoucherHandler) QueryVoucherOfShop(c *gin.Context) {

	shopId := c.Param("shopId")
	if shopId == "" {
		c.JSON(http.StatusBadRequest, result.Fail("输入有误"))
		return
	}
	id, _ := strconv.ParseInt(shopId, 10, 64)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vos, err := h.svc.FindVoucher(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.OkData(vos))

}
