package service

import (
	"errors"
	"go-admin/app/admin/models"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"gorm.io/gorm"

	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type SysMetaData struct {
	service.Service
}

// Get 获取SysDept对象
func (e *SysMetaData) Get(d *dto.SysMetaDataGetReq, model *models.SysMetaData) error {
	var err error
	var data models.SysMetaData

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysDept对象
func (e *SysMetaData) Insert(c *dto.SysMetaDataInsertReq) error {
	var err error
	var data models.SysMetaData
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	metaDataPath := pkg.IntToString(data.MetaDataId) + "/"
	if data.ParentId != 0 {
		var metaDataP models.SysMetaData
		tx.First(&metaDataP, data.ParentId)
		metaDataPath = metaDataP.MetaDataPath + metaDataPath
	} else {
		metaDataPath = "/0/" + metaDataPath
	}
	var mp = map[string]string{}
	mp["meta_data_path"] = metaDataPath
	if err := tx.Model(&data).Update("meta_data_path", metaDataPath).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysDept对象
func (e *SysMetaData) Update(c *dto.SysMetaDataUpdateReq) error {
	var err error
	var model = models.SysMetaData{}
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx.First(&model, c.GetId())
	c.Generate(&model)

	metaDataPath := pkg.IntToString(model.MetaDataId) + "/"
	if model.ParentId != 0 {
		var metaDataP models.SysMetaData
		tx.First(&metaDataP, model.ParentId)
		metaDataPath = metaDataP.MetaDataPath + metaDataPath
	} else {
		metaDataPath = "/0/" + metaDataPath
	}
	model.MetaDataPath = metaDataPath
	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateSysMetaData error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysDept
func (e *SysMetaData) Remove(d *dto.SysMetaDataDeleteReq) error {
	var err error
	var data models.SysMetaData

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetSysDeptList 获取组织数据
func (e *SysMetaData) getList(c *dto.SysMetaDataGetPageReq, list *[]models.SysMetaData) error {
	var err error
	var data models.SysMetaData

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// SetMetaDataPage 设置metadata页面数据
func (e *SysMetaData) SetMetaDataPage(c *dto.SysMetaDataGetPageReq) (m []models.SysMetaData, err error) {
	var list []models.SysMetaData
	err = e.getList(c, &list)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := e.metaDataPageCall(&list, list[i])
		m = append(m, info)
	}
	return
}

func (e *SysMetaData) metaDataPageCall(metaDatatlist *[]models.SysMetaData, menu models.SysMetaData) models.SysMetaData {
	list := *metaDatatlist
	min := make([]models.SysMetaData, 0)
	for j := 0; j < len(list); j++ {
		if menu.MetaDataId != list[j].ParentId {
			continue
		}
		mi := models.SysMetaData{}
		mi.MetaDataId = list[j].MetaDataId
		mi.ParentId = list[j].ParentId
		mi.MetaDataPath = list[j].MetaDataPath
		mi.MetaDataName = list[j].MetaDataName
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []models.SysMetaData{}
		ms := e.metaDataPageCall(metaDatatlist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}
