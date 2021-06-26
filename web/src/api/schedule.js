import service from '@/utils/request'

// @Tags Schedule
// @Summary 创建Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Schedule true "创建Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /schedule/createSchedule [post]
export const createSchedule = (data) => {
  return service({
    url: '/schedule/createSchedule',
    method: 'post',
    data
  })
}

// @Tags Schedule
// @Summary 删除Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Schedule true "删除Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /schedule/deleteSchedule [delete]
export const deleteSchedule = (data) => {
  return service({
    url: '/schedule/deleteSchedule',
    method: 'delete',
    data
  })
}

// @Tags Schedule
// @Summary 删除Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /schedule/deleteSchedule [delete]
export const deleteScheduleByIds = (data) => {
  return service({
    url: '/schedule/deleteScheduleByIds',
    method: 'delete',
    data
  })
}

// @Tags Schedule
// @Summary 更新Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Schedule true "更新Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /schedule/updateSchedule [put]
export const updateSchedule = (data) => {
  return service({
    url: '/schedule/updateSchedule',
    method: 'put',
    data
  })
}

// @Tags Schedule
// @Summary 用id查询Schedule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Schedule true "用id查询Schedule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /schedule/findSchedule [get]
export const findSchedule = (params) => {
  return service({
    url: '/schedule/findSchedule',
    method: 'get',
    params
  })
}

// @Tags Schedule
// @Summary 分页获取Schedule列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Schedule列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /schedule/getScheduleList [get]
export const getScheduleList = (params) => {
  return service({
    url: '/schedule/getScheduleList',
    method: 'get',
    params
  })
}