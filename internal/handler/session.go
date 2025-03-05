package handler

import (
	"github.com/gin-gonic/gin"
	"go-gin/internal/model"
	"gorm.io/gorm"
	"net/http"
)

func GenNewSession(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var convParams model.Conversation

		if err := c.ShouldBindJSON(&convParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 改进后的代码
		tx := db.Begin()
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "事务启动失败"})
			return
		}

		// 使用匿名函数封装事务逻辑便于错误处理
		defer func() {
			if r := recover(); r != nil { // 处理panic
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "系统异常"})
			}
		}()

		//----------**here is wrong**------------------------------

		newSession := model.Conversation{
			ConvType: convParams.ConvType,
			Title:    "new session",
			OwnerID:  convParams.OwnerID,
			Members:  []*model.User{{UserID: convParams.OwnerID}},
			// members no need

		}
		//---------------------------------
		// 创建消息记录（带错误处理）
		if err := tx.Create(&newSession).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "session setup failed: " + err.Error()})
			return
		}

		// 提交事务（带错误检查）
		if err := tx.Commit().Error; err != nil {
			tx.Rollback() // 提交失败时回滚
			c.JSON(http.StatusInternalServerError, gin.H{"error": "事务提交失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"session_id": newSession.ConvID})
	}

}

/*
func GenGetSession(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) { return }
}
*/

func GenGetSessions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 定义请求参数结构体
		type QueryParams struct {
			Page     int    `form:"page,default=1"`  // 分页页码[2](@ref)
			PageSize int    `form:"size,default=20"` // 每页数量[2](@ref)
			UserID   uint64 `form:"user_id"`         // 用户ID过滤条件
		}

		// 2. 参数绑定与验证[1](@ref)
		var params QueryParams
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		//--------------
		var results []model.Conversation // 替换为实际模型
		query := db.Model(&model.Conversation{})

		// 添加过滤条件
		if params.UserID > 0 {
			query = query.Select("conv_id, owner_id, created_at").Where("owner_id = ?", params.UserID).Order("created_at desc")
		}
		//--------------------------------------------
		// 3. 构建数据库查询

		// 4. 执行分页查询[2](@ref)
		var total int64
		err := query.Count(&total).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		offset := (params.Page - 1) * params.PageSize
		if err := query.Offset(offset).Limit(params.PageSize).Find(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// 5. 构建标准化响应[1](@ref)
		c.JSON(http.StatusOK, gin.H{
			"list":  results,
			"total": total,
			"page":  params.Page,
			"size":  params.PageSize,
		})
	}
}
