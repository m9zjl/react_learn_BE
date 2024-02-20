package entity

import (
	"time"
)

type User struct {
	ID          int       `gorm:"primary_key " json:"id"`
	Email       string    `json:"email"`
	Nickname    string    `json:"nickname"`
	Passwd      string    `json:"passwd"`
	Ext         string    `json:"ext"`
	GmtCreate   time.Time `json:"gmt_create"`
	GmtModified time.Time `json:"gmt_modified"`
}

func (User) TableName() string {
	return "user"
}
