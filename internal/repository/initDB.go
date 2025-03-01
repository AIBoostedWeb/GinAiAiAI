// internal/repository/database.go
package repository

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"time"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.LoadDatabaseConfig()

	dsn := buildDSN(cfg)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewGormLogger(), // 自定义日志集成Zap
	})

	if err != nil {
		zap.L().Fatal("数据库连接失败",
			zap.String("dsn", maskPassword(dsn)),
			zap.Error(err))
	}

	setupConnectionPool(db, cfg)

	DB = db
}

func buildDSN(cfg *config.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Charset,
		cfg.ParseTime)
}

func setupConnectionPool(db *gorm.DB, cfg *config.DatabaseConfig) {
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Fatal("获取数据库实例失败", zap.Error(err))
	}

	// 连接池设置
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	if duration, err := time.ParseDuration(cfg.ConnMaxLifetime); err == nil {
		sqlDB.SetConnMaxLifetime(duration)
	}
}

// 敏感信息脱敏
func maskPassword(dsn string) string {
	return regexp.MustCompile(`:([^@]+)@`).ReplaceAllString(dsn, ":****@")
}
