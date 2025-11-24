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

type BlogHandler struct {
	svc *application.BlogService
}

func NewBlogHandler(svc *application.BlogService) *BlogHandler {
	return &BlogHandler{svc: svc}
}

func (h *BlogHandler) SaveBlog(c *gin.Context) {
	var dto application.BlogDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Fail("绑定错误"))
		return
	}

	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Fail("无权限"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	blogId, err := h.svc.Create(ctx, &dto, id.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.OkData(blogId))

}

func (h *BlogHandler) LikeBlog(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, result.Fail("获取blogid失败"))
		return
	}
	blogId, _ := strconv.ParseInt(id, 10, 64)
	userId, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Fail("无权限"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.LikeBlog(ctx, userId.(int64), blogId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Ok())

}

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

func (h *BlogHandler) QueryHotBlog(c *gin.Context) {

	page := c.DefaultQuery("current", "1")
	pageSize := int(15)
	pageInt, _ := strconv.Atoi(page)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vos, err := h.svc.FindPageHot(ctx, pageInt, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.OkData(vos))

}

func (h *BlogHandler) QueryBlogById(c *gin.Context) {

	id := c.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindById(ctx, idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.OkData(vo))
}

func (h *BlogHandler) QueryTopLikes(c *gin.Context) {

	id := c.Param("id")
	blogId, _ := strconv.ParseInt(id, 10, 64)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	vos, err := h.svc.TopLikes(ctx, blogId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.OkData(vos))
}
