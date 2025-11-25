//go:build wireinject

package initialize

import (
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/application"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/cache"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/database"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/persistence"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/interfaces/http"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/interfaces/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func Init() (*gin.Engine, error) {
	wire.Build(
		NewRouter,

		http.NewShopTypeHandler,
		http.NewUserHandler,
		http.NewBlogHandler,
		http.NewShopHandler,
		http.NewFollowHandler,

		middleware.AuthMiddleware,

		application.NewShopTypeService,
		application.NewUserService,
		application.NewSessionService,
		application.NewBlogService,
		application.NewShopService,
		application.NewLikeService,
		application.NewFollowService,

		persistence.NewDefaultShopTypeRepository,
		persistence.NewDefaultUserRepository,
		persistence.NewDefaultBlogRepository,
		persistence.NewDefaultShopRepository,
		persistence.NewCachedShopRepository,
		persistence.NewDefaultFollowRepository,

		cache.NewSessionRepository,
		cache.NewLikeRepository,
		cache.NewFollowCache,
		cache.NewGeoCache,

		database.NewMysqlDatabase,
		database.NewRedisDatabase,
	)

	return nil, nil
}
