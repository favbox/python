// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOriginalURL = "original_url"

// OriginalURL mapped from table <original_url>
type OriginalURL struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Type      string         `gorm:"column:type;not null;comment:1:doc,2:img" json:"type"` // 1:doc,2:img
	URL       string         `gorm:"column:url;not null" json:"url"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName OriginalURL's table name
func (*OriginalURL) TableName() string {
	return TableNameOriginalURL
}
