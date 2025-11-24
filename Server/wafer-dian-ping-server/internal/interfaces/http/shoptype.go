package http

import (
	"context"
	"net/http"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/application"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type ShopTypeHandler struct {
	svc *application.ShopTypeService
}

func NewShopTypeHandler(svc *application.ShopTypeService) *ShopTypeHandler {
	return &ShopTypeHandler{svc: svc}
}

func (h *ShopTypeHandler) QueryTypeList(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	vos, err := h.svc.FindShopTypeList(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.OkData(vos))
}
