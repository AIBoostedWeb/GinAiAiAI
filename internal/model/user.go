package model

type User struct {
	ID           uint         `json:"id" gorm:"primaryKey" binding:"-"`
	Username     string       `json:"username" binding:"required"`
	PasswordHash string       `json:"password_hash" binding:"required"`
	Love         LoveStruct   `json:"love" gorm:"embedded"`
	Commit       CommitStruct `json:"commit" gorm:"embedded"`
	MessageID    string       `json:"message_id"`
}
type LimitedUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoveStruct struct {
}
type CommitStruct struct {
}
