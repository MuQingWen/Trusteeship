package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Class struct {
	global.GVA_MODEL
	Cno      string `json:"cno" form:"cno" gorm:"column:cno;comment:班级序号;type:varchar(255);"`
	Cname    string `json:"cname" form:"cname" gorm:"column:cname;comment:班级科目;type:varchar(255);"`
	Chours   int    `json:"chours" form:"chours" gorm:"column:chours;comment:班级课时;type:int;size:10;"`
	Ccycle   string `json:"ccycle" form:"ccycle" gorm:"column:ccycle;comment:班级周期;type:varchar(255);"`
	Ctime    string `json:"ctime" form:"ctime" gorm:"column:ctime;comment:上课时间;type:varchar(255);"`
	Cplace   string `json:"cplace" form:"cplace" gorm:"column:cplace;comment:上课地点;type:varchar(255);"`
	Tno      string `json:"tno" form:"tno" gorm:"column:tno;comment:教师序号;type:varchar(255);"`
	Tgrade   int    `json:"tgrade" form:"tgrade" gorm:"column:tgrade;comment:教师等级;type:int;size:10;"`
	Caccount int    `json:"caccount" form:"caccount" gorm:"column:caccount;comment:班级学费;type:int;size:10;"`
	Cneed    int    `json:"cneed" form:"cneed" gorm:"column:cneed;comment:班级招生人数;type:int;size:10;"`
	Cadd     int    `json:"cadd" form:"cadd" gorm:"column:cadd;comment:班级报名人数;type:int;size:10;"`
}

func (Class) TableName() string {
	return "class"
}
