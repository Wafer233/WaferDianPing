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

type FollowHandler struct {
	svc *application.FollowService
}

func NewFollowHandler(svc *application.FollowService) *FollowHandler {
	return &FollowHandler{svc: svc}
}

func (h *FollowHandler) IsFollow(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, result.Fail("id不存在"))
		return
	}
	userId, _ := strconv.ParseInt(id, 10, 64)

	curId, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Fail("未授权"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ok, err := h.svc.IsFollow(ctx, userId, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.OkData(ok))

}

func (h *FollowHandler) Follow(c *gin.Context) {
	id := c.Param("id")
	isFollow := c.Param("isFollow")
	if id == "" || isFollow == "" {
		c.JSON(http.StatusBadRequest, result.Fail("id不存在"))
		return
	}

	isFol, _ := strconv.ParseBool(isFollow)
	userId, _ := strconv.ParseInt(id, 10, 64)

	curId, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Fail("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := h.svc.Follow(ctx, userId, curId.(int64), isFol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.Ok())

}

func (h *FollowHandler) CommonFollow(c *gin.Context) {

	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	curId, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Fail("未授权"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	vos, err := h.svc.CommonFollow(ctx, userId, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.OkData(vos))
}
