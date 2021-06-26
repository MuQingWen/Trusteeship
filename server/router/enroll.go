package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEnrollRouter(Router *gin.RouterGroup) {
	EnrollRouter := Router.Group("enroll").Use(middleware.OperationRecord())
	{
		EnrollRouter.POST("createEnroll", v1.CreateEnroll)   // 新建Enroll
		EnrollRouter.DELETE("deleteEnroll", v1.DeleteEnroll) // 删除Enroll
		EnrollRouter.DELETE("deleteEnrollByIds", v1.DeleteEnrollByIds) // 批量删除Enroll
		EnrollRouter.PUT("updateEnroll", v1.UpdateEnroll)    // 更新Enroll
		EnrollRouter.GET("findEnroll", v1.FindEnroll)        // 根据ID获取Enroll
		EnrollRouter.GET("getEnrollList", v1.GetEnrollList)  // 获取Enroll列表
	}
}
