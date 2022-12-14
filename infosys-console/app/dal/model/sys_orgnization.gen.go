// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysOrgnization = "sys_orgnization"

// SysOrgnization mapped from table <sys_orgnization>
type SysOrgnization struct {
	ID            int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name          string         `gorm:"column:name" json:"name"`
	Pid           int32          `gorm:"column:pid" json:"pid"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ContactName   string         `gorm:"column:contact_name" json:"contact_name"`
	ContactPhone  string         `gorm:"column:contact_phone" json:"contact_phone"`
	ContactAddres string         `gorm:"column:contact_addres" json:"contact_addres"`
	Ext1          string         `gorm:"column:ext1" json:"ext1"`
	Ext2          string         `gorm:"column:ext2" json:"ext2"`
	Ext3          string         `gorm:"column:ext3" json:"ext3"`
	Ext4          string         `gorm:"column:ext4" json:"ext4"`
	Ext5          string         `gorm:"column:ext5" json:"ext5"`
	Ext6          string         `gorm:"column:ext6" json:"ext6"`
	Ext7          string         `gorm:"column:ext7" json:"ext7"`
	Ext8          string         `gorm:"column:ext8" json:"ext8"`
	Ext9          string         `gorm:"column:ext9" json:"ext9"`
	Ext10         string         `gorm:"column:ext10" json:"ext10"`
	PName         string         `gorm:"column:p_name" json:"p_name"` // 权限名称
	PKey          string         `gorm:"column:p_key" json:"p_key"`   // 权限标识
	Path          string         `gorm:"column:path" json:"path"`
}

// TableName SysOrgnization's table name
func (*SysOrgnization) TableName() string {
	return TableNameSysOrgnization
}
