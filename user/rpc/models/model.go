package models

import (
	"errors"
	"gorm.io/gorm"
	"microservice/user/rpc/internal/utils"
)

type User struct {
	gorm.Model
	Username      string `gorm:"type:varchar(64);not null;unique;index:uniqueIndex"`
	Password      string `gorm:"type:varchar(64);not null"`
	FollowCount   int64  `gorm:"not null"`
	FollowerCount int64  `gorm:"not null"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if len(u.Password) < 6 {
		return errors.New("请设置6位以上的密码")
	}
	u.Password = utils.Password(u.Password)
	return nil
}
