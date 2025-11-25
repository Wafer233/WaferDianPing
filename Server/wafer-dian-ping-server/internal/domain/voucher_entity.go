package domain

import "time"

type Voucher struct {
	Id          int64      `gorm:"primaryKey;column:id;autoIncrement"`
	ShopId      int64      `gorm:"column:shop_id"`
	Title       string     `gorm:"column:title"`
	SubTitle    string     `gorm:"column:sub_title"`
	Rules       string     `gorm:"column:rules"`
	PayValue    int64      `gorm:"column:pay_value"`
	ActualValue int64      `gorm:"column:actual_value"`
	Type        int        `gorm:"column:type"`
	Status      int        `gorm:"column:status"`
	CreateTime  *time.Time `gorm:"column:create_time"`
	UpdateTime  *time.Time `gorm:"column:update_time"`
}

func (Voucher) TableName() string {
	return "tb_voucher"
}

type SecKillVoucher struct {
	VoucherId  int64      `gorm:"column:voucher_id"`
	Stock      int        `gorm:"column:stock"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
	BeginTime  *time.Time `gorm:"column:begin_time"`
	EndTime    *time.Time `gorm:"column:end_time"`
}

func (SecKillVoucher) TableName() string {
	return "tb_seckill_voucher"
}
