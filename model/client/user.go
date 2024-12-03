package client

import (
	"ring/global"
	"time"
)

// User C端用户
type User struct {
	global.BaseModel
	Username     string    `json:"username" gorm:"type:varchar(32);comment:用户名;uniqueIndex;"`
	Password     string    `json:"-" gorm:"type:varchar(64);comment:密码"`
	Nickname     string    `json:"nickname" gorm:"type:varchar(32);comment:昵称;"`
	Avatar       string    `json:"avatar" gorm:"type:varchar(512);comment:头像URL"`
	Email        string    `json:"email" gorm:"type:varchar(64);comment:邮箱;uniqueIndex;"`
	Phone        string    `json:"phone" gorm:"type:varchar(32);comment:手机号;uniqueIndex;"`
	RegistryTime time.Time `json:"registryTime" gorm:"comment:注册时间"`
}
