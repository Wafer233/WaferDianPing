package domain

import "time"

type VoucherOrder struct {
	Id         int64      `gorm:"column:id"`
	UserId     int64      `gorm:"column:user_id"`
	VoucherId  int64      `gorm:"column:voucher_id"`
	PayType    int        `gorm:"column:pay_type"`
	Status     int        `gorm:"column:status"`
	CreateTime *time.Time `gorm:"column:create_time"`
	PayTime    *time.Time `gorm:"column:pay_time"`
	UseTime    *time.Time `gorm:"column:use_time"`
	RefundTime *time.Time `gorm:"column:refund_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
}

func (VoucherOrder) TableName() string {
	return "tb_voucher_order"
}
