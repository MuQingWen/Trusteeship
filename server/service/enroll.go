package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateEnroll
//@description: 创建Enroll记录
//@param: enroll model.Enroll
//@return: err error

func CreateEnroll(enroll model.Enroll) (err error) {
	err = global.GVA_DB.Create(&enroll).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteEnroll
//@description: 删除Enroll记录
//@param: enroll model.Enroll
//@return: err error

func DeleteEnroll(enroll model.Enroll) (err error) {
	err = global.GVA_DB.Delete(&enroll).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteEnrollByIds
//@description: 批量删除Enroll记录
//@param: ids request.IdsReq
//@return: err error

func DeleteEnrollByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Enroll{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateEnroll
//@description: 更新Enroll记录
//@param: enroll *model.Enroll
//@return: err error

func UpdateEnroll(enroll model.Enroll) (err error) {
	err = global.GVA_DB.Save(&enroll).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetEnroll
//@description: 根据id获取Enroll记录
//@param: id uint
//@return: err error, enroll model.Enroll

func GetEnroll(id uint) (err error, enroll model.Enroll) {
	err = global.GVA_DB.Where("id = ?", id).First(&enroll).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetEnrollInfoList
//@description: 分页获取Enroll记录
//@param: info request.EnrollSearch
//@return: err error, list interface{}, total int64

func GetEnrollInfoList(info request.EnrollSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Enroll{})
    var enrolls []model.Enroll
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Cno != "" {
        db = db.Where("`cno` = ?",info.Cno)
    }
    if info.Cname != "" {
        db = db.Where("`cname` = ?",info.Cname)
    }
    if info.Ccycle != "" {
        db = db.Where("`ccycle` = ?",info.Ccycle)
    }
    if info.Ctime != "" {
        db = db.Where("`ctime` = ?",info.Ctime)
    }
    if info.Caccount != 0 {
        db = db.Where("`caccount` = ?",info.Caccount)
    }
    if info.Sname != "" {
        db = db.Where("`sname` = ?",info.Sname)
    }
    if info.Saccount != "" {
        db = db.Where("`saccount` = ?",info.Saccount)
    }
    if info.Sphone != "" {
        db = db.Where("`sphone` = ?",info.Sphone)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&enrolls).Error
	return err, enrolls, total
}