package request

import "gin-vue-admin/model"

type EnrollSearch struct{
    model.Enroll
    PageInfo
}