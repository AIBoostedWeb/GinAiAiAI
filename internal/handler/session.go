package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenSession(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) { return }
}

func GenGetSession(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) { return }
}
