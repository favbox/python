// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameWeixinMp = "weixin_mp"

// WeixinMp mapped from table <weixin_mp>
type WeixinMp struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID             int64     `gorm:"column:uid;not null" json:"uid"`
	AuthorID        int64     `gorm:"column:author_id;not null" json:"author_id"`
	Biz             string    `gorm:"column:biz;not null" json:"biz"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Avatar          string    `gorm:"column:avatar;not null" json:"avatar"`
	LastPublishTime time.Time `gorm:"column:last_publish_time" json:"last_publish_time"`
	CreateTime      time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName WeixinMp's table name
func (*WeixinMp) TableName() string {
	return TableNameWeixinMp
}