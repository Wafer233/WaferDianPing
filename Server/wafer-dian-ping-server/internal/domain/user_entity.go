package domain

import "time"

type User struct {
	Id         int64      `gorm:"column:id;primaryKey;type:bigint,autoIncrement"`
	Phone      string     `gorm:"column:phone;varchar(11)"`
	NickName   string     `gorm:"column:nick_name;varchar(128)"`
	Icon       string     `gorm:"column:icon;varchar(255)"`
	CreateTime *time.Time `gorm:"column:create_time;type:timestamp"`
	UpdateTime *time.Time `gorm:"column:update_time;type:timestamp"`
}

func (User) TableName() string {
	return "tb_user"
}

type UserInfo struct {
	UserId     int64      `gorm:"column:user_id;type:bigint"`
	City       string     `gorm:"column:city;varchar(64)"`
	Introduce  string     `gorm:"column:introduce"`
	Fans       int        `gorm:"column:fans"`
	Followee   int        `gorm:"column:followee"`
	Gender     int        `gorm:"column:gender"`
	Birthday   *time.Time `gorm:"column:birthday;type:date"`
	Credits    int        `gorm:"column:credits"`
	Level      int        `gorm:"column:level"`
	CreateTime *time.Time `gorm:"column:create_time;type:timestamp"`
	UpdateTime *time.Time `gorm:"column:update_time;type:timestamp"`
}

func (UserInfo) TableName() string {
	return "tb_user_info"
}
