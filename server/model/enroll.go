// 自动生成模板Enroll
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Enroll struct {
      global.GVA_MODEL
      Cno  string `json:"cno" form:"cno" gorm:"column:cno;comment:班级序号;type:varchar(255);"`
      Cname  string `json:"cname" form:"cname" gorm:"column:cname;comment:班级名称;type:varchar(255);"`
      Ccycle  string `json:"ccycle" form:"ccycle" gorm:"column:ccycle;comment:班级周期;type:varchar(255);"`
      Tname  string `json:"tname" form:"tname" gorm:"column:tname;comment:教师序号;type:varchar(255);"`
      Ctime  string `json:"ctime" form:"ctime" gorm:"column:ctime;comment:上课时间;type:varchar(255);"`
      Caccount  int `json:"caccount" form:"caccount" gorm:"column:caccount;comment:班级学费;type:int;size:10;"`
      Sname  string `json:"sname" form:"sname" gorm:"column:sname;comment:学生姓名;type:varchar(255);"`
      Saccount  string `json:"saccount" form:"saccount" gorm:"column:saccount;comment:学生缴费;type:varchar(255);"`
      Sphone  string `json:"sphone" form:"sphone" gorm:"column:sphone;comment:学生电话;type:varchar(255);"`
}


func (Enroll) TableName() string {
  return "enroll"
}

