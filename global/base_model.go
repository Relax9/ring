package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"time"
)

type AuditKey struct{}

// BaseModel 基础模型属性
type BaseModel struct {
	Id        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `gorm:"comment:创建时间;autoCreateTime"`
	CreatedBy uint           `gorm:"comment:创建人;"`
	UpdatedAt time.Time      `gorm:"comment:最后修改时间;autoUpdateTime;"`
	UpdatedBy uint           `gorm:"comment:修改人"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateAudit(db *gorm.DB) {
	createBy := db.Statement.Schema.LookUpField("CreatedBy")
	if createBy == nil {
		return
	}

	context, ok := db.Statement.Context.(*gin.Context)
	var userId string
	if !ok {
		return
	}

	userId, ok = context.Request.Context().Value(AuditKey{}).(string)
	if !ok {
		userId = "System"
	}

	data := db.Statement.ReflectValue
	if data.Kind() == reflect.Slice {
		for i := 0; i < data.Len(); i++ {
			element := data.Index(i)
			field := element.FieldByName("CreatedBy")
			if field.IsValid() && field.CanSet() {
				conv, err := strconv.Atoi(userId)
				if err != nil {
					conv = 0
				}
				field.SetUint(uint64(conv))
			}
		}
		return
	}

	db.Statement.SetColumn("CreatedBy", userId)
}

func UpdateAudit(db *gorm.DB) {
	db.Statement.Schema.LookUpField("UpdatedBy")
}
