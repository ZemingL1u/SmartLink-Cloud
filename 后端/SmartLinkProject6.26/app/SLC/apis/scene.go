// apis包定义了SmartLinkProject的API接口
package apis

import (
	"SmartLinkProject/app/SLC/models" // 导入项目中的设备模型
	"net/http"                        // 导入net包，用于处理HTTP相关功能

	"github.com/gin-gonic/gin"         // 导入Gin框架，用于处理HTTP请求
	"github.com/gin-gonic/gin/binding" // 导入Gin框架的binding功能，用于数据绑定

	"github.com/go-admin-team/go-admin-core/sdk/api" // 导入go-admin-core的api模块
	// 导入go-admin-core的response模块，虽然未使用，但可能用于响应处理
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"SmartLinkProject/app/SLC/service"     // 导入项目中的设备服务层
	"SmartLinkProject/app/SLC/service/dto" // 导入项目中的设备数据传输对象（DTO）

	"SmartLinkProject/common/actions" // 导入项目中的权限操作
)

type Scene struct {
	api.Api
}

func (e Scene) GetPage(c *gin.Context) {
	// 实例化设备服务
	s := service.Scene{}
	// 初始化设备分页请求数据传输对象（DTO）
	req := dto.SceneGetPageReq{}
	// 绑定请求数据到req DTO，创建ORM和上下文，并将服务实例赋值给s
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		// 记录错误日志并返回500错误响应
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 从Gin上下文中获取当前用户的权限信息
	p := actions.GetPermissionFromContext(c)

	// 准备用于接收查询结果的切片和计数器变量
	list := make([]models.Scene, 0)
	var count int64

	// 调用服务层的GetPage方法来获取设备分页列表和总记录数
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		// 如果服务层返回错误，返回500错误响应
		e.Error(500, err, "查询失败")
		return
	}
	// 返回分页成功的响应，包括设备列表、当前页码、每页大小和总记录数
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// 其他方法（如Get、Insert、Update、Delete、UpdateStatus、GetProfile）遵循类似的模式：
// 1. 实例化服务层对象
// 2. 初始化请求数据传输对象（DTO）
// 3. 绑定请求数据到DTO并创建服务层上下文
// 4. 执行服务层方法，并处理错误
// 5. 根据服务层执行结果返回相应的HTTP响应

func (e Scene) Get(c *gin.Context) {
	s := service.Scene{}
	req := dto.SceneById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Scene
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

func (e Scene) Insert(c *gin.Context) {
	s := service.Scene{}
	req := dto.SceneInsertReq{}
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

	e.OK(req.GetId(), "创建成功")
}

func (e Scene) Update(c *gin.Context) {
	s := service.Scene{}
	req := dto.SceneUpdateReq{}
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
	e.OK(req.GetId(), "更新成功")
}

func (e Scene) Delete(c *gin.Context) {
	s := service.Scene{}
	req := dto.SceneById{}
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

func (e Scene) GetSceneIdsByLocId(c *gin.Context) {
	s := service.Scene{}
	req := dto.DLocJoin{}
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

	SceneIds, err := s.GetSceneIdsByLocId(&req)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(SceneIds, "查询成功")
}

func (e Scene) GetPageBySceneIds(c *gin.Context) {
	s := service.Scene{}
	req := dto.SceneGetPageReqBySceneIds{}

	// 绑定请求数据到req DTO，创建ORM和上下文，并将服务实例赋值给s
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

	// 准备用于接收查询结果的切片和计数器变量
	list := make([]models.Scene, 0)
	var count int64

	// 从Gin上下文中获取当前用户的权限信息
	p := actions.GetPermissionFromContext(c)

	// 调用服务层的GetPage方法来获取设备分页列表和总记录数
	// 注意：这里的实现取决于service包中Scene服务的GetPage方法是否支持通过ID列表进行查询
	err = s.GetPageBySceneIds(req.SceneIds, &req, p, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	// 返回分页成功的响应，包括设备列表、当前页码、每页大小和总记录数
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}
