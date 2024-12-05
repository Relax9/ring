package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"ring/global"
	"ring/model/client"
	"ring/model/system"
)

func InitMySQL() {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/ring?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Callback().Create().Before("gorm:before_create").Register("create_audit", global.CreateAudit)
	if err != nil {
		panic("failed to register create audit")
	}
	err = db.Callback().Update().Before("gorm:before_update").Register("update_audit", global.UpdateAudit)
	if err != nil {
		panic("failed to register update audit")
	}

	err = db.AutoMigrate(&system.SysUser{}, &client.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	global.DB = db
}
