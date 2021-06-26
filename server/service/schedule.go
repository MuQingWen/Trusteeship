package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSchedule
//@description: 创建Schedule记录
//@param: schedule model.Schedule
//@return: err error

func CreateSchedule(schedule model.Schedule) (err error) {
	err = global.GVA_DB.Create(&schedule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSchedule
//@description: 删除Schedule记录
//@param: schedule model.Schedule
//@return: err error

func DeleteSchedule(schedule model.Schedule) (err error) {
	err = global.GVA_DB.Delete(&schedule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteScheduleByIds
//@description: 批量删除Schedule记录
//@param: ids request.IdsReq
//@return: err error

func DeleteScheduleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Schedule{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSchedule
//@description: 更新Schedule记录
//@param: schedule *model.Schedule
//@return: err error

func UpdateSchedule(schedule model.Schedule) (err error) {
	err = global.GVA_DB.Save(&schedule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSchedule
//@description: 根据id获取Schedule记录
//@param: id uint
//@return: err error, schedule model.Schedule

func GetSchedule(id uint) (err error, schedule model.Schedule) {
	err = global.GVA_DB.Where("id = ?", id).First(&schedule).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetScheduleInfoList
//@description: 分页获取Schedule记录
//@param: info request.ScheduleSearch
//@return: err error, list interface{}, total int64

func GetScheduleInfoList(info request.ScheduleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Schedule{})
    var schedules []model.Schedule
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Cno != "" {
        db = db.Where("`cno` = ?",info.Cno)
    }
    if info.Cname != "" {
        db = db.Where("`cname` = ?",info.Cname)
    }
    if info.Tno != "" {
        db = db.Where("`tno` = ?",info.Tno)
    }
    if info.Sname != "" {
        db = db.Where("`sname` = ?",info.Sname)
    }
    if info.Ctime != "" {
        db = db.Where("`ctime` = ?",info.Ctime)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&schedules).Error
	return err, schedules, total
}