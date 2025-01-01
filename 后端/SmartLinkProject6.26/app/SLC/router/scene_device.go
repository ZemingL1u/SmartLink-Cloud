package router

import (
	"SmartLinkProject/app/SLC/apis"
	"SmartLinkProject/common/actions"
	"SmartLinkProject/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSceneDeviceRouter)
}

// 需认证的路由代码
func registerSceneDeviceRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SceneDevice{}
	r := v1.Group("/sceneDevice").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.POST("/get", api.Get)
		r.POST("", api.Insert)
		r.DELETE("", api.Delete)
	}
}
