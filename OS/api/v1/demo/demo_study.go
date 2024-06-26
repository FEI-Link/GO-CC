package demo

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type StudyReq struct {
	g.Meta `path:"/study/list" tags:"详情查询" method:"get" summary:"详情查询"`
	UserId int `json:"userId"   description:"这个ID和个人相关联"`
}
type StudyRes struct {
	g.Meta `mime:"application/json"`
	List   gdb.Result `json:"list"`
}
type StudyAddReq struct {
	g.Meta   `path:"/study/add" tags:"详情查询" method:"post" summary:"详情查询"`
	Id       uint        `json:"id"       description:""`
	UserId   int         `json:"userId"   description:""`
	Study    string      `json:"study"    description:"学习经历"`
	StudyDe  string      `json:"studyDe"  description:"学习经历详情"`
	Createat *gtime.Time `json:"createat" description:""`
	Delete   *gtime.Time `json:"delete"   description:""`
}
type StudyAddRes struct {
}
type StudyEditReq struct {
	g.Meta  `path:"/study/edit" tags:"详情查询" method:"put" summary:"详情查询"`
	Id      uint   `json:"id"       description:""`
	UserId  int    `json:"userId"   description:""`
	Study   string `json:"study"    description:"学习经历"`
	StudyDe string `json:"studyDe"  description:"学习经历详情"`
}
type StudyEditRes struct {
}
type StudyDeleteReq struct {
	g.Meta `path:"/study/delete" tags:"详情查询" method:"delete" summary:"详情查询"`
	Id     uint `json:"id"       description:""`
}
type StudyDeleteRes struct {
}
