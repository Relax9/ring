package global

import "time"

// BaseModel 基础模型属性
type BaseModel struct {
	Id        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
