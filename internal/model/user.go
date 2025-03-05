package model

import "time"

type User struct {
	UserID       uint64 `json:"user_id" gorm:"primaryKey;column:user_id" binding:"-"`
	Username     string `json:"username" gorm:"size:50;not null;unique" binding:"required"`
	PasswordHash string `json:"password_hash" gorm:"not null" binding:"required"`
	//Love         LoveStruct   `json:"love" gorm:"embedded;embeddedPrefix:love_"`
	//Commit       CommitStruct `json:"commit" gorm:"embedded;embeddedPrefix:commit_"`
	Messages  []*Message `json:"messages" gorm:"foreignKey:SenderID;references:UserID"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type LoveStruct struct {
}
type CommitStruct struct {
}
