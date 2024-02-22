package entity

import "time"

type Article struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"type:varchar(256)" json:"title"`
	Desc        string    `gorm:"type:text" json:"desc"`
	Img         string    `gorm:"type:varchar(1024)" json:"img"`
	Uid         int       `json:"uid"`
	UserId      int       `json:"user_id"`
	Category    string    `gorm:"type:varchar(64)" json:"category"`
	GmtCreate   time.Time `json:"gmt_create"`
	GmtModified time.Time `json:"gmt_modified"`
}

func (Article) TableName() string {
	return "article"
}
