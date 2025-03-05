package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin/pkg/util"
	"net/http"
	"strings"
)

// JWTAuth JWT 鉴权中间件（支持黑名单校验）
func JWTAuth(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从 Header 中提取 Token
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			var err error
			authHeader, err = c.Cookie("token")
			if err != nil || authHeader == "" {
				abortWithError(c, http.StatusUnauthorized, "缺失认证令牌")
				return
			}

		}

		// 2. 验证 Token 格式
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			abortWithError(c, http.StatusUnauthorized, "令牌格式错误，应为: Bearer <token>")
			return
		}
		tokenString := tokenParts[1]

		// 3. 解析并验证 Token
		claims, err := util.ParseToken(secret, tokenString)
		if err != nil {
			handleParseError(c, err)
			return
		}

		// 5. 存储用户信息到上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("token_exp", claims.ExpiresAt.Unix())

		c.Next()
	}
}

func handleParseError(c *gin.Context, err error) {
	switch {
	case strings.Contains(err.Error(), "expired"):
		abortWithError(c, http.StatusUnauthorized, "令牌已过期")
	case strings.Contains(err.Error(), "signature"):
		abortWithError(c, http.StatusUnauthorized, "签名验证失败")
	default:
		abortWithError(c, http.StatusUnauthorized, "无效令牌")
	}
}

// 终止请求并返回错误
func abortWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
		"error":   true,
	})
}
