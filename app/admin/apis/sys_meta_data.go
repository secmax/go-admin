package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysMetaData struct {
	api.Api
}

// GetPage
// @Summary 分页源数据列表数据
// @Description 分页列表
// @Tags 源数据表
// @Param deptName query string false "metaDataName"
// @Param deptId query string false "metaDataId"
// @Param position query string false "position"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/metadata [get]
// @Security Bearer
func (e SysMetaData) GetPage(c *gin.Context) {
	s := service.SysMetaData{}
	req := dto.SysMetaDataGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	list := make([]models.SysMetaData, 0)
	list, err = s.SetMetaDataPage(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list, "查询成功")
}

// Get
// @Summary 获取部门数据
// @Description 获取JSON
// @Tags 部门
// @Param deptId path string false "deptId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security Bearer
func (e SysMetaData) Get(c *gin.Context) {
	s := service.SysMetaData{}
	req := dto.SysMetaDataGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysMetaData

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert 添加元数据
// @Summary 添加元数据
// @Description 获取JSON
// @Tags 元数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysMetaDataInsertReq true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/metaData [post]
// @Security Bearer
func (e SysMetaData) Insert(c *gin.Context) {
	s := service.SysMetaData{}
	req := dto.SysMetaDataInsertReq{}
	fmt.Print("=============:", &c)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update
// @Summary 修改
// @Description 获取JSON
// @Tags
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysMetaDataUpdateReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/metaData/{metaDataId} [put]
// @Security Bearer
func (e SysMetaData) Update(c *gin.Context) {
	s := service.SysMetaData{}
	req := dto.SysMetaDataUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除
// @Description 删除数据
// @Tags
// @Param data body dto.SysMetaDataDeleteReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/metaData [delete]
// @Security Bearer
func (e SysMetaData) Delete(c *gin.Context) {
	s := service.SysMetaData{}
	req := dto.SysMetaDataDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}
