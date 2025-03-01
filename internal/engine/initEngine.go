package engine

import (
	"github.com/gin-gonic/gin"
	"go-gin/internal/pkg/middleware"
	"go.uber.org/zap"
)

func InitMainHandler(runMode string, logger *zap.Logger) *gin.Engine {
	gin.SetMode(runMode)
	r := gin.New()

	r.Use(
		middleware.ZapLogger(logger),
		middleware.ZapRecovery(logger),
	)

	return r
}
