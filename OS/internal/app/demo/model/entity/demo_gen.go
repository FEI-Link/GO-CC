// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DemoGen is the golang structure for table demo_gen.
type DemoGen struct {
	Id         uint        `json:"id"         description:""`
	DemoName   string      `json:"demoName"   description:"姓名"`
	DemoAge    uint        `json:"demoAge"    description:"年龄"`
	Classes    string      `json:"classes"    description:"班级"`
	DemoBorn   *gtime.Time `json:"demoBorn"   description:"出生年月"`
	DemoGender uint        `json:"demoGender" description:"性别"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建日期"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"修改日期"`
	DeletedAt  *gtime.Time `json:"deletedAt"  description:"删除日期"`
	CreatedBy  uint64      `json:"createdBy"  description:"创建人"`
	UpdatedBy  uint64      `json:"updatedBy"  description:"修改人"`
	DemoStatus int         `json:"demoStatus" description:"状态"`
	DemoCate   string      `json:"demoCate"   description:"分类"`
	DemoThumb  string      `json:"demoThumb"  description:"头像"`
	DemoPhoto  string      `json:"demoPhoto"  description:"相册"`
	DemoInfo   string      `json:"demoInfo"   description:"个人描述"`
	DemoFile   string      `json:"demoFile"   description:"相关附件"`
	ClassesTwo string      `json:"classesTwo" description:"班级二"`
}
