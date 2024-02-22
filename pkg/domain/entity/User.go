package entity

import (
	"time"
)

type User struct {
	ID          int       `gorm:"primary_key " json:"id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Passwd      string    `json:"passwd,omitempty"`
	Ext         string    `json:"ext"`
	GmtCreate   time.Time `json:"gmt_create"`
	GmtModified time.Time `json:"gmt_modified"`
}

func (User) TableName() string {
	return "user"
}
