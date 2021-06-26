package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitScheduleRouter(Router *gin.RouterGroup) {
	ScheduleRouter := Router.Group("schedule").Use(middleware.OperationRecord())
	{
		ScheduleRouter.POST("createSchedule", v1.CreateSchedule)   // 新建Schedule
		ScheduleRouter.DELETE("deleteSchedule", v1.DeleteSchedule) // 删除Schedule
		ScheduleRouter.DELETE("deleteScheduleByIds", v1.DeleteScheduleByIds) // 批量删除Schedule
		ScheduleRouter.PUT("updateSchedule", v1.UpdateSchedule)    // 更新Schedule
		ScheduleRouter.GET("findSchedule", v1.FindSchedule)        // 根据ID获取Schedule
		ScheduleRouter.GET("getScheduleList", v1.GetScheduleList)  // 获取Schedule列表
	}
}
