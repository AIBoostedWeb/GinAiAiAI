package handler

import (
	"github.com/gin-gonic/gin"
	"go-gin/internal/config"
	"go-gin/internal/model"
	"go-gin/internal/service"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GenHandleMessage(db *gorm.DB, aiCfg *config.AIConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		var messageParams model.Message
		var response string
		if err := c.ShouldBindJSON(&messageParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//------------ AIGC ----------------
		history, _ := getMessageInPages(db, c, messageParams.SenderID, messageParams.ConvID, 10, 1)
		chatContext := append(*history, messageParams)

		requestPrompt, i := service.AskAiWithMessage(aiCfg, &chatContext, service.PROMPT)
		if i != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": i.Error()})
			return
		}
		var err error
		var midResponse string
		midResponse, err = service.AskAIWithStr(aiCfg, requestPrompt, service.GENERATE)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err = service.AskAIWithStr(aiCfg, "根据"+response+"检查ai生成内容:\n\n"+response+"\n\n如果有问题就改正，然后不分析，将bash脚本传出，并去掉多余的符号。", service.CHECK)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//(*history)[len(*history)-1].ResponseContent = response
		//----------------------------

		//----------if no conv_id and receiver_id raise error
		//------------------ tx begin --------------------------------------
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

		timeNow := time.Now()
		msg := model.Message{
			SenderID:        messageParams.SenderID,
			Content:         messageParams.Content,
			MsgType:         messageParams.MsgType,
			ConvID:          messageParams.ConvID,
			ReceiverType:    messageParams.ReceiverType,
			SendTime:        timeNow,
			ResponseContent: response,
		}

		// 群聊消息验证（带错误检查）
		if messageParams.ConvID != 0 {
			var count int64

			if err := tx.Table("conversation_users").
				Where("conversation_id = ? AND user_id = ?", messageParams.ConvID, messageParams.SenderID).
				Count(&count).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "成员验证失败"})
				return
			}
			if count == 0 {
				tx.Rollback() // 显式回滚
				c.JSON(http.StatusForbidden, gin.H{"error": "sender not in conversation"})
				return
			}
		}

		// 创建消息记录（带错误处理）
		if err := tx.Create(&msg).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存消息失败: " + err.Error()})
			return
		}

		// 提交事务（带错误检查）
		if err := tx.Commit().Error; err != nil {
			tx.Rollback() // 提交失败时回滚
			c.JSON(http.StatusInternalServerError, gin.H{"error": "事务提交失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message_id": msg.MessageID, "question": messageParams.Content, "prompt": requestPrompt, "midRes": midResponse, "message": response})
	}
}

func GenPullMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 定义请求参数结构体
		type QueryParams struct {
			Page     int    `form:"page,default=1" json:"page" binding:"required"`       // 分页页码[2](@ref)
			PageSize int    `form:"size,default=20" json:"page_size" binding:"required"` // 每页数量[2](@ref)
			UserID   uint64 `form:"user_id" json:"user_id" binding:"required"`           // 用户ID过滤条件
			ConvID   uint64 `form:"conv_id" json:"conv_id" binding:"required"`
		}

		// 2. 参数绑定与验证[1](@ref)
		var params QueryParams
		if err := c.ShouldBindQuery(&params); err != nil {
			if err := c.ShouldBindJSON(&params); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		// 3. 构建数据库查询
		var results []model.Message // 替换为实际模型
		query := db.Model(&model.Message{})

		// 添加过滤条件
		if params.UserID > 0 {
			query = query.Where("sender_id = ?", params.UserID)
		}
		if params.ConvID > 0 { // 新增ConvID过滤
			query = query.Where("conv_id = ?", params.ConvID)
		}

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

func getMessageInPages(db *gorm.DB, c *gin.Context, userId uint64, convId uint64, pageSize int, page int) (*[]model.Message, int64) {
	var results []model.Message // 替换为实际模型
	query := db.Model(&model.Message{})

	// 添加过滤条件
	if userId > 0 {
		query = query.Where("sender_id = ?", userId)
	}
	if convId > 0 { // 新增ConvID过滤
		query = query.Where("conv_id = ?", convId)
	}

	// 4. 执行分页查询[2](@ref)
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return nil, -1
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return nil, -1
	}
	return &results, total
}
