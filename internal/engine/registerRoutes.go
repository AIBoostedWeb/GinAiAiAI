package engine

import (
	"github.com/gin-gonic/gin"
	"go-gin/internal/config"
	"go-gin/internal/handler"
	"go-gin/pkg/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(cfg *config.ServiceConfig, serverConfig *config.ServerConfig, r *gin.Engine, db *gorm.DB) {

	public := r.Group("/")
	{

		public.POST("/login", handler.GenLogin(db, cfg.JWTSecret, serverConfig))
		public.POST("/register", handler.GenRegister(db))
	}

	// 受保护路由（应用中间件）
	protected := r.Group("/api/v1")
	protected.Use(middleware.JWTAuth([]byte(cfg.JWTSecret))) // 注入JWT服务（可为nil）
	{

		user := protected.Group("/user")
		// user.GET("/", handler.GetUserProfile) // 动态路由参数设计[7](@ref)
		// user.PUT("/", handler.UpdateProfile)
		user.POST("/logout", handler.GenLogout(db, serverConfig.Host))

		session := protected.Group("/session")
		session.POST("/", handler.GenNewSession(db))
		session.GET("/", handler.GenGetSessions(db))

		sms := protected.Group("/message")
		sms.POST("/", handler.GenHandleMessage(db, &serverConfig.AI))
		sms.GET("/", handler.GenPullMessage(db))

	}

}
