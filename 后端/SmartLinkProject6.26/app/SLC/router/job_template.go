package router

import (
	"SmartLinkProject/app/SLC/apis"
	"SmartLinkProject/common/actions"
	"SmartLinkProject/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerJobTemplateRouter)
}

// 需认证的路由代码
func registerJobTemplateRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.JobTemplate{}
	r := v1.Group("/jobTemplate").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("/pageList", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}

}
