package router

import (
	"SmartLinkProject/app/SLC/apis"
	"SmartLinkProject/common/actions"
	"SmartLinkProject/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSceneRouter)
}

// 需认证的路由代码
func registerSceneRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Scene{}
	r := v1.Group("/scene").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.GET("/getSceneIdByLocId", api.GetSceneIdsByLocId)
		r.GET("/getListBySceneIds", api.GetPageBySceneIds)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}
}
