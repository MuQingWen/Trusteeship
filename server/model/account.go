// 自动生成模板Account
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Account struct {
      global.GVA_MODEL
      Cno  string `json:"cno" form:"cno" gorm:"column:cno;comment:班级序号;type:varchar(255);"`
      Cname  string `json:"cname" form:"cname" gorm:"column:cname;comment:中文名称;type:varchar(255);"`
      Caccount  int `json:"caccount" form:"caccount" gorm:"column:caccount;comment:班级学费;type:int;size:10;"`
      Saccount  int `json:"saccount" form:"saccount" gorm:"column:saccount;comment:学生缴费;type:int;size:10;"`
      Sphone  string `json:"sphone" form:"sphone" gorm:"column:sphone;comment:学生电话;type:varchar(255);"`
}


func (Account) TableName() string {
  return "account"
}

