package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"ring/api/system"
	"ring/config"
	"ring/global"
)

func main() {

	config.InitMySQL()
	config.InitRedis()
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.WithValue(c, global.AuditKey{}, "123456"))
		c.Next()
	})

	engine.GET("/api/users", system.CreateSystemUser)
	if gin.Mode() == "debug" {
		engine.Use(gin.Logger())
	}
	err := engine.Run(":8080")
	if err != nil {
		panic("fail to start server" + err.Error())
	}
}

//
//package main
//
//import (
//	"context"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type User struct {
//	gorm.Model
//	Name  string
//	Email string
//}
//type contextKey struct{}
//
//func main() {
//	// 初始化Gin路由器实例
//	r := gin.Default()
//
//	// 初始化Gorm数据库连接
//	dsn := "root:123456@tcp(127.0.0.1:3306)/ring?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//	db.Callback().Create().Before("gorm:create").Register("update_created_at", updateCreated)
//	// 自动迁移模型结构体到数据库表
//	db.AutoMigrate(&User{})
//
//	// 创建处理程序函数
//	createUser := func(c *gin.Context) {
//		user := User{
//			Name: "leo" + time.Now().Format("2006-01-02 15:04:05"),
//		}
//		db.WithContext(c).Create(&user)
//		c.JSON(200, user)
//	}
//
//	// 自定义中间件，设置example 变量为123
//	r.Use(func(c *gin.Context) {
//		c.Request = c.Request.WithContext(context.WithValue(c, contextKey{}, "123@qq.com"))
//		c.Next()
//	})
//	// 在路由器实例中注册处理程序函数
//
//	r.GET("/users", createUser)
//
//	// 启动Gin服务器，监听HTTP请求
//	r.Run(":8080")
//}
//
//func updateCreated(db *gorm.DB) {
//	ginCtx, ok := db.Statement.Context.(*gin.Context)
//	if !ok {
//		return
//	}
//	email, ok := ginCtx.Request.Context().Value(contextKey{}).(string)
//	if ok {
//		if _, ok := db.Statement.Schema.FieldsByName["Email"]; ok {
//			db.Statement.SetColumn("Email", email)
//		}
//	}
//}
