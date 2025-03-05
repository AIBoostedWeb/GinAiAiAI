package model

import (
	"gorm.io/gorm"
	"time"
)

type Conversation struct {
	ConvID   uint64  `json:"conv_id" gorm:"primaryKey;autoIncrement;comment:会话ID"`
	ConvType int8    `json:"conv_type" gorm:"not null;default:1;index;comment:1-单聊 2-群聊" binding:"required"`
	Title    string  `json:"title" gorm:"size:255;index;comment:会话标题"`
	OwnerID  uint64  `json:"owner_id" gorm:"not null;index;comment:拥有者ID" binding:"required"`
	Members  []*User `json:"members" gorm:"many2many:conversation_users;foreignKey:ConvID;joinForeignKey:ConversationID;References:UserID;joinReferences:UserID"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;comment:软删除时间"`
}
