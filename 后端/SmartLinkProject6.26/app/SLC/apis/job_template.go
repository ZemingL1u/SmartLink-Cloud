package apis

import (
	"SmartLinkProject/app/SLC/models" // 导入项目中的设备模型
	// "encoding/json"
	"net/http" // 导入net包，用于处理HTTP相关功能
	// "strconv"

	"github.com/gin-gonic/gin"         // 导入Gin框架，用于处理HTTP请求
	"github.com/gin-gonic/gin/binding" // 导入Gin框架的binding功能，用于数据绑定

	"github.com/go-admin-team/go-admin-core/sdk/api" // 导入go-admin-core的api模块
	// 导入go-admin-core的response模块，虽然未使用，但可能用于响应处理
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"SmartLinkProject/app/SLC/service"     // 导入项目中的设备服务层
	"SmartLinkProject/app/SLC/service/dto" // 导入项目中的设备数据传输对象（DTO）

	"SmartLinkProject/common/actions" // 导入项目中的权限操作
	// "net/url"
)

type JobTemplate struct {
	api.Api
}

func (e JobTemplate) Get(c *gin.Context) {
	s := service.JobTemplate{}
	req := dto.JobTemplateGetReq{}
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
	// 准备用于接收查询结果的切片和计数器变量
	data := make([]models.JobTemplate, 0)
	var count int64
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.GetByTypeId(&req, p, &data, &count)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	// 返回分页成功的响应，包括任务模板列表、当前页码、每页大小和总记录数
	e.PageOK(data, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e JobTemplate) Insert(c *gin.Context) {
	s := service.JobTemplate{}
	req := dto.JobTemplateInsertReq{}
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
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "创建成功")
}

func (e JobTemplate) Update(c *gin.Context) {
	s := service.JobTemplate{}
	req := dto.JobTemplateUpdateReq{}
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

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "更新成功")
}

func (e JobTemplate) Delete(c *gin.Context) {
	s := service.JobTemplate{}
	req := dto.TemplateById{}
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

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "删除成功")
}
