package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gin/internal/config"
	"go-gin/internal/model"
	"go-gin/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

// GenRegister 定义 Handler 生成函数（正确注入 db）
func GenRegister(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 绑定请求数据
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求格式"})
			return
		}

		if len(user.PasswordHash) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "密码长度至少6位"})
			return
		}

		// 检查用户是否已存在（使用注入的 db 实例）
		var existingUser model.User
		result := db.Where("username = ?", user.Username).First(&existingUser)
		if result.Error == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
			return
		} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 处理其他数据库错误
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
			return
		}

		// 密码哈希处理
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
			return
		}
		user.PasswordHash = string(hashedPassword)

		// 保存到数据库（使用注入的 db 实例）
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败"})
			return
		}

		// 返回成功响应（RESTful 规范）
		c.JSON(http.StatusCreated, gin.H{
			"code":    http.StatusCreated,
			"message": "注册成功",
			"data":    gin.H{"user_id": user.UserID}, // 返回用户ID更符合规范
		})
	}
}

func GenLogin(db *gorm.DB, jwtSecret string, cfg *config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		var userInfo model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := db.Where("Username = ?", user.Username).First(&userInfo).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(userInfo.PasswordHash), []byte(user.PasswordHash)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
			return
		}

		// 4. 生成 JWT Token（示例）
		token, err := util.GenerateToken([]byte(jwtSecret), userInfo.UserID, userInfo.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "user_id": userInfo.UserID})
		c.SetCookie("token", token, 2400000, "/", cfg.Host, false, true)

	}
}

func GenLogout(db *gorm.DB, domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", domain, false, true)

		// 可选：清理服务端会话（如Redis中的Token记录）
		// redisClient.Del(c, "user_session:"+userID)

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "成功注销",
		})
	}
}
