package models

import (
	"time"

	"github.com/flyrory/go-web-blog/pkg/types"
)

// BaseModel 模型基类
type BaseModel struct {
	ID        uint64
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
