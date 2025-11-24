package http

import (
	"context"
	"net/http"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/application"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	svc *application.BlogService
}

func NewBlogHandler(svc *application.BlogService) *BlogHandler {
	return &BlogHandler{svc: svc}
}

func (h *BlogHandler) SaveBlog(c *gin.Context) {
}

func (h *BlogHandler) LikeBlog(c *gin.Context) {}

func (h *BlogHandler) QueryMyBlog(c *gin.Context) {
	curId, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Fail("未授权"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	vos, err := h.svc.FindBlogByUserId(ctx, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.OkData(vos))

}

func (h *BlogHandler) QueryHotBlog(c *gin.Context) {}
