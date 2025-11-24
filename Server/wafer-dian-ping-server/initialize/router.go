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
) *gin.Engine {

	r := gin.Default()
	r = http.NewShopTypeRouter(r, shty)
	r = http.NewUserRouter(r, user, auth)
	r = http.NewBlogRouter(r, blog, auth)
	return r

}
