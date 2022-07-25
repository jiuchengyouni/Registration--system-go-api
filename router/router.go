package router

import (
	"Registration-system/controller"
	"Registration-system/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	//处理跨域
	r.Use(middleware.Cors())
	r.POST("/student/apply",controller.Apply)
	r.POST("/student/login",controller.Login)
	//加了token
	//r.POST("/student/order",controller.Order,middleware.AuthMiddleware())
	//r.POST("/student/wait",controller.Wait,middleware.AuthMiddleware())
	r.POST("/student/order",controller.Order)
	r.POST("/student/wait",controller.Wait)
	r.POST("/admin/advance/:department",controller.Advance)
	r.POST("/admin/delay/:department",controller.Delay)
	r.POST("/admin/delete/:department",controller.Delete)
	r.POST("/admin/line/:department",controller.Line)
	return r
}