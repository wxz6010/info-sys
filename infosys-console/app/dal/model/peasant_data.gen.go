// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePeasantDatum = "peasant_data"

// PeasantDatum mapped from table <peasant_data>
type PeasantDatum struct {
	ID        int32          `gorm:"column:id;primaryKey" json:"id"`
	UserID    int64          `gorm:"column:user_id" json:"user_id"`
	Name      string         `gorm:"column:name" json:"name"`
	Gender    []uint8        `gorm:"column:gender" json:"gender"`
	Address   string         `gorm:"column:address" json:"address"`
	Idcard    string         `gorm:"column:idcard" json:"idcard"`
	CreateAt  time.Time      `gorm:"column:create_at" json:"create_at"`
	UpdateAt  time.Time      `gorm:"column:update_at" json:"update_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName PeasantDatum's table name
func (*PeasantDatum) TableName() string {
	return TableNamePeasantDatum
}