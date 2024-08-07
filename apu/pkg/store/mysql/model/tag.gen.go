// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTag = "tag"

// Tag mapped from table <tag>
type Tag struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Type       string    `gorm:"column:type;not null" json:"type"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
