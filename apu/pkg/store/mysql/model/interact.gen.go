// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInteract = "interact"

// Interact mapped from table <interact>
type Interact struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NoteID         int64     `gorm:"column:note_id;not null" json:"note_id"`
	Day            int32     `gorm:"column:day;not null" json:"day"`
	ReadCount      int32     `gorm:"column:read_count;not null;comment:阅读量" json:"read_count"`           // 阅读量
	LikedCount     int32     `gorm:"column:liked_count;not null;comment:点赞量" json:"liked_count"`         // 点赞量
	CollectedCount int32     `gorm:"column:collected_count;not null;comment:收藏量" json:"collected_count"` // 收藏量
	CommentCount   int32     `gorm:"column:comment_count;not null;comment:评论量" json:"comment_count"`     // 评论量
	ShareCount     int32     `gorm:"column:share_count;not null;comment:分享量" json:"share_count"`         // 分享量
	CreateTime     time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName Interact's table name
func (*Interact) TableName() string {
	return TableNameInteract
}
