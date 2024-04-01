package demo

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DetailsReq struct {
	g.Meta   `path:"/details/list" tags:"详情查询" method:"get" summary:"详情查询"`
	UserName string `json:"userName"   description:"用户名"`
}
type DetailsRes struct {
	g.Meta `mime:"application/json"`
	List   gdb.Result `json:"list"`
}
type DetailsAddReq struct {
	g.Meta   `path:"/details/add" tags:"新增详情" method:"get" summary:"新增详情"`
	UserName string      `json:"userName"   description:"用户名"`
	Position string      `json:"position"   description:"职位信息"`
	Bumen    string      `json:"bumen"      description:"所属部门"`
	Ruzhiat  *gtime.Time `json:"ruzhiat"    description:"入职时间"`
	State    string      `json:"state"      description:"目前在职状态"`
}
type DetailsAddRes struct {
}
