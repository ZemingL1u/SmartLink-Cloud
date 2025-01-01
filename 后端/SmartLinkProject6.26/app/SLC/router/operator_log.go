package router

import (
	"SmartLinkProject/app/SLC/apis"
	"SmartLinkProject/common/actions"
	"SmartLinkProject/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerOperatorLogRouter)
}

// 需认证的路由代码
func registerOperatorLogRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.OperatorLog{}
	r := v1.Group("/operatorLog").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("/pageList", api.GetPage)
		r.POST("", api.Insert)
		r.DELETE("", api.Delete)
	}

}
