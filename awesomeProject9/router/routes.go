package router

import (
	"github.com/gin-gonic/gin"
	"oceanlearn.teach/ginessential/controller"
	"oceanlearn.teach/ginessential/middleware"
)

//CollectRouter
func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register",controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.GET("api/auth/info",middleware.AuthMiddleware(),controller.Info)
	return r
}
