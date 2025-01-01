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

type SceneDevice struct {
	api.Api
}

// 其他方法（如Get、Insert、Update、Delete、UpdateStatus、GetProfile）遵循类似的模式：
// 1. 实例化服务层对象
// 2. 初始化请求数据传输对象（DTO）
// 3. 绑定请求数据到DTO并创建服务层上下文
// 4. 执行服务层方法，并处理错误
// 5. 根据服务层执行结果返回相应的HTTP响应

func (e SceneDevice) Get(c *gin.Context) {
	s := service.SceneDevice{}
	req := dto.SceneDeviceGetReq{}
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
	print("-------------------")
	print(req.SceneId)
	print("-------------------")

	var object models.SceneDevice
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

func (e SceneDevice) Insert(c *gin.Context) {
	s := service.SceneDevice{}
	req := dto.SceneDeviceInsertReq{}
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

func (e SceneDevice) Delete(c *gin.Context) {
	s := service.SceneDevice{}
	req := dto.SceneDeviceRemoveReq{}
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
