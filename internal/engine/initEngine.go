package engine

import (
	"github.com/gin-gonic/gin"
	middleware2 "go-gin/pkg/middleware"
	"go.uber.org/zap"
)

func InitMainHandler(runMode string, logger *zap.Logger) *gin.Engine {
	gin.SetMode(runMode)

	r := gin.New()

	r.Use(
		middleware2.ZapLogger(logger),
		middleware2.ZapRecovery(logger),
	)

	return r
}
