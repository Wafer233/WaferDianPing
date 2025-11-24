package domain

import "time"

type ShopType struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string     `gorm:"column:name"`
	Icon       string     `gorm:"column:icon"`
	Sort       int        `gorm:"column:sort"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
}

func (ShopType) TableName() string {
	return "tb_shop_type"
}
