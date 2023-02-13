package model

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	Phone         string `gorm:"column:phone"`
	Email         string `gorm:"column:email"`
	Identity      string
	ClientIP      string
	DeviceInfo    string
	IsLogout      bool
	LoginTime     uint64
	LogoutTime    uint64
	HeartBeatTine uint64
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
