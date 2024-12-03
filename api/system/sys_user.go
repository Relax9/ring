package system

import (
	"github.com/gin-gonic/gin"
	"ring/global"
	"ring/model/system"
	"time"
)

func CreateSystemUser(ctx *gin.Context) {
	user := system.SysUser{
		Username: "leo" + time.Now().Format("2006-01-02 15:04:05"),
	}
	global.DB.WithContext(ctx).Create(&user)
}
