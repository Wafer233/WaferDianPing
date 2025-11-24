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

type UserHandler struct {
	svc *application.UserService
}

func NewUserHandler(svc *application.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) SendCode(c *gin.Context) {}

func (h *UserHandler) Login(c *gin.Context) {

	dto := application.LoginDTO{}

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail(err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sessionId, err := h.svc.LoginService(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	if sessionId == "" {
		c.JSON(http.StatusInternalServerError, result.Fail("无法获取sessionId"))
		return
	}

	//设置cookie
	c.SetCookie("session_id", sessionId, 3600, "/", "localhost", false, false)

	c.JSON(http.StatusOK, result.Ok())
}

func (h *UserHandler) Logout(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, result.Fail(err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = h.svc.LogoutService(ctx, sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.Ok())
}

func (h *UserHandler) Me(c *gin.Context) {

	curId, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Fail("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	vo, err := h.svc.FindUser(ctx, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.OkData(vo))
}

func (h *UserHandler) Info(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, result.Fail("id空"))
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	vo, err := h.svc.FindUserInfo(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.OkData(vo))
}
