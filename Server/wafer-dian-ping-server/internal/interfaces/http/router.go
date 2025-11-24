package http

import (
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, h *UserHandler, auth gin.HandlerFunc) *gin.Engine {

	user := r.Group("/user/")

	user.POST("login", h.Login)
	user.Use(auth)
	user.GET("me", h.Me)
	user.POST("logout", h.Logout)
	user.GET("info/:id", h.Info)

	return r
}

func NewBlogRouter(r *gin.Engine, h *BlogHandler, auth gin.HandlerFunc) *gin.Engine {

	blog := r.Group("/blog/")
	blog.Use(auth)
	blog.POST("", h.SaveBlog)
	blog.PUT("like/:id", h.LikeBlog)
	blog.GET("likes/:id", h.QueryTopLikes)
	blog.GET("of/me", h.QueryMyBlog)
	blog.GET("hot", h.QueryHotBlog)
	blog.GET(":id", h.QueryBlogById)

	return r

}

func NewShopRouter(r *gin.Engine, h *ShopHandler, auth gin.HandlerFunc) *gin.Engine {

	shop := r.Group("/shop/")
	shop.Use(auth)
	shop.GET("of/name", h.QueryShopByName)
	shop.GET(":id", h.QueryShopById)
	shop.GET("of/type", h.QueryShopByType)

	return r

}

func NewShopTypeRouter(r *gin.Engine, h *ShopTypeHandler) *gin.Engine {

	shopType := r.Group("/shop-type/")

	shopType.GET("list", h.QueryTypeList)
	return r
}
