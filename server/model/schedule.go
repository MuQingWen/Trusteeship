// 自动生成模板Schedule
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Schedule struct {
      global.GVA_MODEL
      Cno  string `json:"cno" form:"cno" gorm:"column:cno;comment:班级序号;type:varchar(255);"`
      Cname  string `json:"cname" form:"cname" gorm:"column:cname;comment:班级科目;type:varchar(255);"`
      Tno  string `json:"tno" form:"tno" gorm:"column:tno;comment:教师序号;type:varchar(255);"`
      Sname  string `json:"sname" form:"sname" gorm:"column:sname;comment:学生姓名;type:varchar(255);"`
      Ctime  string `json:"ctime" form:"ctime" gorm:"column:ctime;comment:上课时间;type:varchar(255);"`
}


func (Schedule) TableName() string {
  return "schedule"
}

