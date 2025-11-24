package initialize

import (
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/interfaces/http"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	shty *http.ShopTypeHandler,
	user *http.UserHandler,
	blog *http.BlogHandler,
	auth gin.HandlerFunc,
	shops *http.ShopHandler,
) *gin.Engine {

	r := gin.Default()

	r = http.NewUserRouter(r, user, auth)
	r = http.NewBlogRouter(r, blog, auth)
	r = http.NewShopRouter(r, shops, auth)
	r = http.NewShopTypeRouter(r, shty)
	return r

}
