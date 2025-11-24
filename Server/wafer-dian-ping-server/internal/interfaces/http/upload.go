package http

import "github.com/gin-gonic/gin"

type UploadHandler struct {
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (h *UploadHandler) UploadImage(c *gin.Context) {}

func (h *UploadHandler) DeleteBlogImg(c *gin.Context) {}

func (h *UploadHandler) CreateNewFileName(c *gin.Context) {}
