package demo

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ReviewReq struct {
	g.Meta `path:"/review/list" tags:"详情查询" method:"get" summary:"详情查询"`
	UserId int `json:"userId"   description:"这个ID和个人相关联"`
}
type ReviewRes struct {
	g.Meta `mime:"application/json"`
	List   gdb.Result `json:"list"`
}
type ReviewAddReq struct {
	g.Meta   `path:"/review/add" tags:"详情查询" method:"post" summary:"详情查询"`
	Id       uint        `json:"id"       description:""`
	UserId   int         `json:"userId"   description:""`
	Study    string      `json:"study"    description:"学习经历"`
	StudyDe  string      `json:"studyDe"  description:"学习经历详情"`
	Createat *gtime.Time `json:"createat" description:""`
	Delete   *gtime.Time `json:"delete"   description:""`
}
type ReviewAddRes struct {
}
type ReviewEditReq struct {
	g.Meta  `path:"/review/edit" tags:"详情查询" method:"put" summary:"详情查询"`
	Id      uint   `json:"id"       description:""`
	UserId  int    `json:"userId"   description:""`
	Study   string `json:"study"    description:"学习经历"`
	StudyDe string `json:"studyDe"  description:"学习经历详情"`
}
type ReviewEditRes struct {
}
type ReviewDeleteReq struct {
	g.Meta `path:"/review/delete" tags:"详情查询" method:"delete" summary:"详情查询"`
	Id     uint `json:"id"       description:""`
}
type ReviewDeleteRes struct {
}
