package system

import "ring/global"

// SysUser 系统用户
type SysUser struct {
	global.BaseModel
	Username string `json:"username" gorm:"type:varchar(32);comment:用户名;uniqueIndex;"`
	Password string `json:"-" gorm:"type:varchar(64);comment:密码"`
	Nickname string `json:"nickname" gorm:"type:varchar(32);comment:昵称"`
	Avatar   string `json:"avatar" gorm:"type:varchar(512);comment:头像URL"`
	Phone    string `json:"phone" gorm:"type:varchar(32);comment:手机号;uniqueIndex;"`
	Email    string `json:"email" gorm:"type:varchar(64);comment:邮箱地址;uniqueIndex;"`
	Enable   bool   `json:"enable" gorm:"comment:是否启用"`
	IsAdmin  bool   `json:"isAdmin" gorm:"comment:是否管理员"`
}
