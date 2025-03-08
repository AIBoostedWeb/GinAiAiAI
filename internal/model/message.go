package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	MessageID       uint64         `json:"message_id" gorm:"primaryKey;autoIncrement;comment:消息唯一ID"`
	SenderID        uint64         `json:"sender_id"  gorm:"not null;index;comment:发送者ID"`
	Content         string         `json:"content" binding:"required" gorm:"type:text;not null;check:content <> '';comment:消息内容"`
	ResponseContent string         `json:"response_content" binding:"-" gorm:"type:text;check:content <> '';comment:ai内容"`
	MsgType         *int8          `json:"msg_type" binding:"required" gorm:"default:1;comment:消息类型（1-文本 2-图片）"`
	ReceiverType    *int8          `json:"receiver_type" binding:"required" gorm:"default:1;comment:消息类型（1-to-session 2-to-user）"`
	SendTime        time.Time      `json:"send_time" gorm:"type:timestamp(3);index:idx_convid_time,priority:2,sort:desc;comment:发送时间"`
	ConvID          uint64         `json:"conv_id" gorm:"not null;index:idx_convid;comment:会话ID"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index;comment:软删除时间"`
	// ReceiverID uint64         `json:"receiver_id" binding:"required" gorm:"not null;index:idx_receiver_time,priority:1;index:idx_convid_time,priority:1;comment:接收者ID"`
}

type MessageOnly struct {
	Content         string `json:"content" binding:"required" gorm:"type:text;not null;check:content <> '';comment:消息内容"`
	ResponseContent string `json:"response_content" binding:"-" gorm:"type:text;default:'';check:content <> '';comment:ai内容"`
	MsgType         *int8  `json:"msg_type" binding:"required" gorm:"default:1;comment:消息类型（1-文本 2-图片）"`
}
