package http

import (
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, h *UserHandler, auth gin.HandlerFunc) *gin.Engine {

	user := r.Group("/user/")

	user.POST("login", h.Login)

	protected := r.Group("/user/")
	protected.Use(auth)

	protected.GET("me", h.Me)
	protected.POST("logout", h.Logout)
	protected.GET("info/:id", h.Info)

	return r
}

func NewBlogRouter(r *gin.Engine, h *BlogHandler, auth gin.HandlerFunc) *gin.Engine {

	blog := r.Group("/blog/")
	blog.Use(auth)
	blog.POST("", h.SaveBlog)
	blog.GET("like/:id", h.LikeBlog)
	blog.GET("of/me", h.QueryMyBlog)
	blog.GET("hot", h.QueryHotBlog)

	return r

}

func NewShopTypeRouter(r *gin.Engine, h *ShopTypeHandler) *gin.Engine {

	shopType := r.Group("/shop-type/")

	shopType.GET("list", h.QueryTypeList)
	return r
}
