package model

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Username      string
	Password      string
	Phone         string
	Email         string
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
