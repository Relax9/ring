package client

import (
	"ring/global"
	"time"
)

// User 消费客户
type User struct {
	global.BaseModel
	Username     string    `json:"username"`
	Password     string    `json:"-"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	RegistryTime time.Time `json:"registryTime"`
}
