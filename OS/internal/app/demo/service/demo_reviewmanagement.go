package service

import (
	"context"
	"gocc/api/v1/demo"
	"gocc/internal/app/demo/dao"
	"gocc/internal/app/demo/model/do"
	"gocc/library/liberr"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var demoRemanService = DemoRemanImpl{}

func DemoReman() IDemoReman {
	return &demoRemanService
}

type DemoRemanImpl struct {
}
type IDemoReman interface {
	DemoRemanList(ctx context.Context, req *demo.RemanReq) (rs gdb.Result, err error)
	DemoRemanAdd(ctx context.Context, req *demo.RemanAddReq) (err error)
	DemoRemanEdit(ctx context.Context, req *demo.RemanEditReq) (err error)
	DemoRemanDetele(ctx context.Context, req *demo.RemanDeleteReq) (err error)
}

func (s *DemoRemanImpl) DemoRemanList(ctx context.Context, req *demo.RemanReq) (rs gdb.Result, err error) {
	rs, err = g.DB().GetAll(ctx, `SELECT s.id,s.user_id,u.user_nickname,s.study,s.study_de,s.createat
	from sys_study s 
	left join sys_user u on s.user_id=u.id
	where s.user_id=?
`, req.UserId)
	return
}
func (s *DemoRemanImpl) DemoRemanAdd(ctx context.Context, req *demo.RemanAddReq) (err error) {
	currentTime := time.Now() // 获取当前时间
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysStudy.Ctx(ctx).Insert(do.SysStudy{
			UserId:   req.UserId,
			Study:    req.Study,
			StudyDe:  req.StudyDe,
			Createat: gtime.New(currentTime),
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}
func (s *DemoRemanImpl) DemoRemanEdit(ctx context.Context, req *demo.RemanEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysStudy.Ctx(ctx).Where(dao.SysStudy.Columns().Id, req.Id).Update(do.SysStudy{
			Study:   req.Study,
			StudyDe: req.StudyDe,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}
func (s *DemoRemanImpl) DemoRemanDetele(ctx context.Context, req *demo.RemanDeleteReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysStudy.Ctx(ctx).Where(dao.SysStudy.Columns().Id, req.Id).Delete()
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}