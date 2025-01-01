package router

import (
	"SmartLinkProject/app/SLC/apis"
	"SmartLinkProject/common/actions"
	"SmartLinkProject/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDeviceRouter)
}

// 需认证的路由代码
func registerDeviceRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Device{}
	r := v1.Group("/device").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.GET("/getDeviceIdByLocId", api.GetDeviceIdsByLocId)
		r.POST("/getlistByDeviceIds", api.GetPageByDeviceIds)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
		r.GET("getTempData", api.GetTemp)
		r.GET("getHumiData", api.GetHumi)
		r.GET("controlBeep", api.ControlBeep)
		r.GET("controlLed", api.ControlLed)
	}

	user := v1.Group("/device").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		user.GET("/profile", api.GetProfile)
		user.PUT("/status", api.UpdateStatus)
	}
}
