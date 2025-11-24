package domain

import "time"

type Shop struct {
	Id         int64      `gorm:"primaryKey;autoIncrement;column:id"`
	Name       string     `gorm:"column:name"`
	TypeId     int64      `gorm:"column:type_id"`
	Images     string     `gorm:"column:images"`
	Area       string     `gorm:"column:area"`
	Address    string     `gorm:"column:address"`
	X          float64    `gorm:"column:x"`
	Y          float64    `gorm:"column:y"`
	AvgPrice   int64      `gorm:"column:avg_price"`
	Sold       int        `gorm:"column:sold"`
	Comments   int        `gorm:"column:comments"`
	Score      int        `gorm:"column:score"`
	OpenHours  string     `gorm:"column:open_hours"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
}

func (Shop) TableName() string {
	return "tb_shop"
}
