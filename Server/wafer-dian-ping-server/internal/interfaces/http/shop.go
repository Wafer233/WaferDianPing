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

type ShopHandler struct {
	svc *application.ShopService
}

func NewShopHandler(svc *application.ShopService) *ShopHandler {
	return &ShopHandler{svc: svc}
}

func (h *ShopHandler) QueryShopById(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, result.Fail("获取id失败"))
		return
	}

	idInt64, _ := strconv.ParseInt(id, 10, 64)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	vo, err := h.svc.FindById(ctx, idInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.OkData(vo))

}

func (h *ShopHandler) SaveShop(c *gin.Context) {}

func (h *ShopHandler) UpdateShop(c *gin.Context) {}

func (h *ShopHandler) QueryShopByType(c *gin.Context) {

	typeId := c.Query("typeId")
	if typeId == "" {
		c.JSON(http.StatusBadRequest, result.Fail("获取id失败"))
		return
	}
	current := c.DefaultQuery("current", "1")

	// 这里是他设置的
	pageSize := int(5)
	page, _ := strconv.Atoi(current)
	typeIdInt, _ := strconv.Atoi(typeId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	vos, err := h.svc.FindPage(ctx, int64(typeIdInt), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.OkData(vos))
}

func (h *ShopHandler) QueryShopByName(c *gin.Context) {}
