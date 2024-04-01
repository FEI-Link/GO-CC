package service

import (
	"context"
	"fmt"
	"gocc/api/v1/demo"
	"gocc/internal/app/demo/dao"
	"gocc/internal/app/demo/model/do"
	"gocc/library/liberr"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var demoDetailsService = DemoDetailsImpl{}

func DemoDetails() IDemoDetails {
	return &demoDetailsService
}

type DemoDetailsImpl struct {
}
type IDemoDetails interface {
	DemoDetails01(ctx context.Context, UserName string) (rs gdb.Result, err error)
	DemoDetailsAdd(ctx context.Context, req *demo.DetailsAddReq) (err error)
}

func (s *DemoDetailsImpl) DemoDetails01(ctx context.Context, UserName string) (rs gdb.Result, err error) {
	rs, err = g.DB().GetAll(ctx, `
	select 
	log_id as xh,
	user_name as xm,
	position as zwxx,
	bumen as ssbm,
	creattime as rzsj,
	state as zzzt
	 from  sys_details where user_id=?`, UserName)
	return
}
func (s *DemoDetailsImpl) DemoDetailsAdd(ctx context.Context, req *demo.DetailsAddReq) (err error) {
	fmt.Print("CESHI测试增加函数是否有效")
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysDetails.Ctx(ctx).Insert(do.SysDetails{
			UserName: req.UserName,
			Position: req.Position,
			Bumen:    req.Bumen,
			Ruzhiat:  req.Ruzhiat,
			State:    req.State,
		})
		liberr.ErrIsNil(ctx, err, "添加黑名单失败")
	})
	return
}
