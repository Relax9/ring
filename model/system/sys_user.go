package system

import "ring/global"

// SysUser 系统用户
type SysUser struct {
	global.BaseModel
	Username string
	Password string
	Nickname string
	Avatar   string
	Phone    string
	Email    string
	Enable   bool
	IsAdmin  bool
}
