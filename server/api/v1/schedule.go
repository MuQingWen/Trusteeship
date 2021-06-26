package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Schedule
// @Summary 创建Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Schedule true "创建Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /schedule/createSchedule [post]
func CreateSchedule(c *gin.Context) {
	var schedule model.Schedule
	_ = c.ShouldBindJSON(&schedule)
	if err := service.CreateSchedule(schedule); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Schedule
// @Summary 删除Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Schedule true "删除Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /schedule/deleteSchedule [delete]
func DeleteSchedule(c *gin.Context) {
	var schedule model.Schedule
	_ = c.ShouldBindJSON(&schedule)
	if err := service.DeleteSchedule(schedule); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Schedule
// @Summary 批量删除Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /schedule/deleteScheduleByIds [delete]
func DeleteScheduleByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteScheduleByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Schedule
// @Summary 更新Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Schedule true "更新Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /schedule/updateSchedule [put]
func UpdateSchedule(c *gin.Context) {
	var schedule model.Schedule
	_ = c.ShouldBindJSON(&schedule)
	if err := service.UpdateSchedule(schedule); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Schedule
// @Summary 用id查询Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Schedule true "用id查询Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /schedule/findSchedule [get]
func FindSchedule(c *gin.Context) {
	var schedule model.Schedule
	_ = c.ShouldBindQuery(&schedule)
	if err, reschedule := service.GetSchedule(schedule.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reschedule": reschedule}, c)
	}
}

// @Tags Schedule
// @Summary 分页获取Schedule列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ScheduleSearch true "分页获取Schedule列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /schedule/getScheduleList [get]
func GetScheduleList(c *gin.Context) {
	var pageInfo request.ScheduleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetScheduleInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
