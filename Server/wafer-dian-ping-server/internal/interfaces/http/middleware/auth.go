package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/application"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(svc *application.SessionService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从cookie中获取token
		sessionId, err := c.Cookie("session_id")
		if sessionId == "" || err != nil {
			c.JSON(http.StatusUnauthorized, result.Fail("未授权"))
			c.Abort()
			return
		}

		userId, err := svc.Get(context.Background(), sessionId)
		if err != nil || userId == "" {
			c.JSON(http.StatusUnauthorized, result.Fail("未授权"))
			c.Abort()
			return
		}

		intId, _ := strconv.Atoi(userId)
		c.Set("id", int64(intId))
		c.Next()
	}
}
