package domain

import "time"

type Follow struct {
	Id           int64      `gorm:"primaryKey;column:id;autoIncrement"`
	UserId       int64      `gorm:"column:user_id"`
	FollowUserId int64      `gorm:"column:follow_user_id"`
	CreateTime   *time.Time `gorm:"column:create_time"`
}

func (Follow) TableName() string {
	return "tb_follow"
}
