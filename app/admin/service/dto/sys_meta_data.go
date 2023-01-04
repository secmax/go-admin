package dto

import (
	"go-admin/app/admin/models"
	common "go-admin/common/models"
)

// SysMetaDataGetPageReq 列表或者搜索使用结构体
type SysMetaDataGetPageReq struct {
	MetaDataId   int    `form:"metaDataId" search:"type:exact;column:meta_data_id;table:sys_meta_data" comment:"id"`       //id
	ParentId     int    `form:"parentId" search:"type:exact;column:parent_id;table:sys_meta_data" comment:"上级部门"`          //上级部门
	MetaDataPath string `form:"metaDataPath" search:"type:exact;column:meta_data_path;table:sys_meta_data" comment:""`     //路径
	MetaDataName string `form:"metaDataName" search:"type:exact;column:meta_data_name;table:sys_meta_data" comment:"部门名称"` //部门名称
	Status       string `form:"status" search:"type:exact;column:status;table:sys_meta_data" comment:"状态"`                 //状态
}

func (m *SysMetaDataGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysMetaDataInsertReq struct {
	MetaDataId   int    `uri:"id" comment:"编码"`                            // 编码
	ParentId     int    `json:"parentId" comment:"上级部门" vd:"?"`            //上级部门
	MetaDataPath string `json:"metaDataPath" comment:""`                   //路径
	MetaDataName string `json:"metaDataName" comment:"部门名称" vd:"len($)>0"` //部门名称
	Status       int    `json:"status" comment:"状态" vd:"$>0"`              //状态
	common.ControlBy
}

func (s *SysMetaDataInsertReq) Generate(model *models.SysMetaData) {
	if s.MetaDataId != 0 {
		model.MetaDataId = s.MetaDataId
	}
	model.MetaDataName = s.MetaDataName
	model.ParentId = s.ParentId
	model.MetaDataPath = s.MetaDataPath
	model.Status = s.Status
}

// GetId 获取数据对应的ID
func (s *SysMetaDataInsertReq) GetId() interface{} {
	return s.MetaDataId
}

type SysMetaDataUpdateReq struct {
	MetaDataId   int    `uri:"id" comment:"编码"`                            // 编码
	ParentId     int    `json:"parentId" comment:"上级部门" vd:"?"`            //上级部门
	MetaDataPath string `json:"metaDataPath" comment:""`                   //路径
	MetaDataName string `json:"metaDataName" comment:"部门名称" vd:"len($)>0"` //部门名称
	Status       int    `json:"status" comment:"状态" vd:"$>0"`              //状态
	common.ControlBy
}

// Generate 结构体数据转化 从 SysDeptControl 至 SysDept 对应的模型
func (s *SysMetaDataUpdateReq) Generate(model *models.SysMetaData) {
	if s.MetaDataId != 0 {
		model.MetaDataId = s.MetaDataId
	}
	model.MetaDataName = s.MetaDataName
	model.ParentId = s.ParentId
	model.MetaDataPath = s.MetaDataPath
	model.Status = s.Status
}

// GetId 获取数据对应的ID
func (s *SysMetaDataUpdateReq) GetId() interface{} {
	return s.MetaDataId
}

type SysMetaDataGetReq struct {
	Id int `uri:"id"`
}

func (s *SysMetaDataGetReq) GetId() interface{} {
	return s.Id
}

type SysMetaDataDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysMetaDataDeleteReq) GetId() interface{} {
	return s.Ids
}

type MetaDataLabel struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []DeptLabel `gorm:"-" json:"children"`
}
