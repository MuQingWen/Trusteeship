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

// @Tags Enroll
// @Summary 创建Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Enroll true "创建Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enroll/createEnroll [post]
func CreateEnroll(c *gin.Context) {
	var enroll model.Enroll
	_ = c.ShouldBindJSON(&enroll)
	if err := service.CreateEnroll(enroll); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Enroll
// @Summary 删除Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Enroll true "删除Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /enroll/deleteEnroll [delete]
func DeleteEnroll(c *gin.Context) {
	var enroll model.Enroll
	_ = c.ShouldBindJSON(&enroll)
	if err := service.DeleteEnroll(enroll); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Enroll
// @Summary 批量删除Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /enroll/deleteEnrollByIds [delete]
func DeleteEnrollByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteEnrollByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Enroll
// @Summary 更新Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Enroll true "更新Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /enroll/updateEnroll [put]
func UpdateEnroll(c *gin.Context) {
	var enroll model.Enroll
	_ = c.ShouldBindJSON(&enroll)
	if err := service.UpdateEnroll(enroll); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Enroll
// @Summary 用id查询Enroll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Enroll true "用id查询Enroll"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /enroll/findEnroll [get]
func FindEnroll(c *gin.Context) {
	var enroll model.Enroll
	_ = c.ShouldBindQuery(&enroll)
	if err, reenroll := service.GetEnroll(enroll.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reenroll": reenroll}, c)
	}
}

// @Tags Enroll
// @Summary 分页获取Enroll列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.EnrollSearch true "分页获取Enroll列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enroll/getEnrollList [get]
func GetEnrollList(c *gin.Context) {
	var pageInfo request.EnrollSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetEnrollInfoList(pageInfo); err != nil {
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
