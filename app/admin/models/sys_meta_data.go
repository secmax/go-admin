package models

import "go-admin/common/models"

type SysMetaData struct {
	MetaDataId   int    `json:"metaDataId" gorm:"primaryKey;autoIncrement;"` // 数据源编码
	ParentId     int    `json:"parentId" gorm:""`                            // 上级关联
	MetaDataPath string `json:"metaDataPath" gorm:"size:255;"`               // 数据源路径
	MetaDataName string `json:"metaDataName"  gorm:"size:128;"`              // 数据源名称
	Status       int    `json:"status" gorm:"size:4;"`                       // 状态
	models.ControlBy
	models.ModelTime
	DataScope string        `json:"dataScope" gorm:"-"`
	Params    string        `json:"params" gorm:"-"`
	Children  []SysMetaData `json:"children" gorm:"-"` // 子表
}

func (SysMetaData) TableName() string {
	return "sys_meta_data"
}

func (e *SysMetaData) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysMetaData) GetId() interface{} {
	return e.MetaDataId
}
