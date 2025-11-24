package domain

import "time"

type Blog struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement"`
	ShopId     int64      `gorm:"column:shop_id"`
	UserId     int64      `gorm:"column:user_id"`
	Title      string     `gorm:"column:title"`
	Images     string     `gorm:"column:images"`
	Content    string     `gorm:"column:content"`
	Liked      int        `gorm:"column:liked"`
	Comments   int        `gorm:"column:comments"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
}

func (Blog) TableName() string {
	return "tb_blog"
}
