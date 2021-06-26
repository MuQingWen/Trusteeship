import service from '@/utils/request'

// @Tags Enroll
// @Summary 创建Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Enroll true "创建Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enroll/createEnroll [post]
export const createEnroll = (data) => {
  return service({
    url: '/enroll/createEnroll',
    method: 'post',
    data
  })
}

// @Tags Enroll
// @Summary 删除Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Enroll true "删除Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /enroll/deleteEnroll [delete]
export const deleteEnroll = (data) => {
  return service({
    url: '/enroll/deleteEnroll',
    method: 'delete',
    data
  })
}

// @Tags Enroll
// @Summary 删除Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /enroll/deleteEnroll [delete]
export const deleteEnrollByIds = (data) => {
  return service({
    url: '/enroll/deleteEnrollByIds',
    method: 'delete',
    data
  })
}

// @Tags Enroll
// @Summary 更新Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Enroll true "更新Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /enroll/updateEnroll [put]
export const updateEnroll = (data) => {
  return service({
    url: '/enroll/updateEnroll',
    method: 'put',
    data
  })
}

// @Tags Enroll
// @Summary 用id查询Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Enroll true "用id查询Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /enroll/findEnroll [get]
export const findEnroll = (params) => {
  return service({
    url: '/enroll/findEnroll',
    method: 'get',
    params
  })
}

// @Tags Enroll
// @Summary 分页获取Enroll列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Enroll列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enroll/getEnrollList [get]
export const getEnrollList = (params) => {
  return service({
    url: '/enroll/getEnrollList',
    method: 'get',
    params
  })
}