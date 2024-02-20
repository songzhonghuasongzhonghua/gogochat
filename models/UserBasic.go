package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `gorm:"type:varchar(20)" json:"name"`
	Password      string `gorm:"type:varchar(255)" json:"password"`
	Phone         string `gorm:"type:varchar(11)" json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `gorm:"type:varchar(100)" json:"email" valid:"email"`
	Identity      string `gorm:"type:varchar(255)" json:"identity"`
	ClientIp      string `gorm:"type:varchar(255)" json:"clientIp"`
	ClientPort    string `gorm:"type:varchar(255)" json:"client_port"`
	LoginTime     uint64 `gorm:"column:login_time;nullable"`
	HeartbeatTime uint64 `gorm:"column:heartbeat_time;nullable"`
	LogoutTime    uint64 `gorm:"column:logout_time;nullable"`
	IsLogout      bool   `gorm:"type:bool" json:"is_logout"`
	DeviceInfo    string `gorm:"type:varchar(255)" json:"device_info"`
}
