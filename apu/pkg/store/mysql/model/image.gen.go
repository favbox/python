// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameImage = "image"

// Image mapped from table <image>
type Image struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID         uint64    `gorm:"column:uid;not null" json:"uid"`
	NoteID      int64     `gorm:"column:note_id;not null" json:"note_id"`
	OriginalURL string    `gorm:"column:original_url;not null" json:"original_url"`
	Width       int32     `gorm:"column:width;not null" json:"width"`
	Height      int32     `gorm:"column:height;not null" json:"height"`
	Sort        int32     `gorm:"column:sort;not null" json:"sort"`
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName Image's table name
func (*Image) TableName() string {
	return TableNameImage
}
