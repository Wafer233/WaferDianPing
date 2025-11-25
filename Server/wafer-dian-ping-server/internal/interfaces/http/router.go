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
	user.GET(":id", h.QueryUserById)

	return r
}

func NewBlogRouter(r *gin.Engine, h *BlogHandler, auth gin.HandlerFunc) *gin.Engine {

	blog := r.Group("/blog/")
	blog.Use(auth)
	blog.POST("", h.SaveBlog)
	blog.GET("of/me", h.QueryMyBlog)
	blog.GET("hot", h.QueryHotBlog)

	blog.PUT("like/:id", h.LikeBlog)
	blog.GET("likes/:id", h.QueryTopLikes)

	blog.GET(":id", h.QueryBlogById)

	return r

}

func NewShopRouter(r *gin.Engine, h *ShopHandler, auth gin.HandlerFunc) *gin.Engine {

	shop := r.Group("/shop/")
	shop.Use(auth)
	shop.GET("of/name", h.QueryShopByName)

	shop.GET("of/type", h.QueryShopByType)

	// 再写 /:id，否则会吃掉所有路径
	shop.GET(":id", h.QueryShopById)

	return r

}

func NewShopTypeRouter(r *gin.Engine, h *ShopTypeHandler) *gin.Engine {

	shopType := r.Group("/shop-type/")

	shopType.GET("list", h.QueryTypeList)
	return r
}

func NewFollowRouter(r *gin.Engine, h *FollowHandler, auth gin.HandlerFunc) *gin.Engine {

	follow := r.Group("/follow/")
	follow.Use(auth)

	follow.GET("or/not/:id", h.IsFollow)
	follow.GET("common/:id", h.CommonFollow)

	follow.PUT(":id/:isFollow", h.Follow)
	return r
}

func NewVoucherRouter(r *gin.Engine, h *VoucherHandler, auth gin.HandlerFunc) *gin.Engine {

	voucher := r.Group("/voucher/")
	voucher.POST("seckill", h.AddSeckillVoucher)
	voucher.POST("", h.AddVoucher)
	voucher.GET("/list/:shopId", h.QueryVoucherOfShop)

	return r
}
