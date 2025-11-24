package domain

import "context"

type UserRepository interface {
	FindUserById(context.Context, int64) (*User, error)
	FindInfoByUserId(context.Context, int64) (*UserInfo, error)
	FindUserByPhone(context.Context, string) (*User, error)
}
