package router

import (
	"Registration-system/controller"
	"Registration-system/docs"
	"Registration-system/middleware"
	swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"

)
func CollectRouter(r *gin.Engine) *gin.Engine {
	//处理跨域
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = ""
	r.Use(middleware.Cors())
	r.POST("/student/apply",controller.Apply)
	r.POST("/student/login",controller.Login)
	//加了token
	//r.POST("/student/order",controller.Order,middleware.AuthMiddleware())
	//r.POST("/student/wait",controller.Wait,middleware.AuthMiddleware())
	r.POST("/student/order",controller.Order)
	r.POST("/student/wait",controller.Wait)
	r.POST("/admin/reset",controller.Reset)
	r.POST("/admin/advance/:department",controller.Advance)
	r.POST("/admin/delay/:department",controller.Delay)
	r.POST("/admin/delete/:department",controller.Delete)
	r.GET("/admin/next/:department",controller.Next)
	r.GET("/admin/line/:department",controller.Line)
	r.GET("/admin/applyLine",controller.ApplyLine)
	return r
}