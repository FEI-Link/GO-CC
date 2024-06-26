package controller

import (
	"context"
	"gocc/api/v1/demo"
	"gocc/internal/app/demo/service"
)

var DemoDetails = demoDetailsController{}

type demoDetailsController struct {
	BaseController
}

func (c *demoDetailsController) DemoDetails01(ctx context.Context, req *demo.DetailsReq) (res *demo.DetailsRes, err error) {
	res = new(demo.DetailsRes)
	res.List, err = service.DemoDetails().DemoDetails01(ctx, req.UserId)
	return
}
func (c *demoDetailsController) DemoDetailsAdd(ctx context.Context, req *demo.DetailsAddReq) (res *demo.DetailsAddRes, err error) {
	err = service.DemoDetails().DemoDetailsAdd(ctx, req)
	return
}
func (c *demoDetailsController) DemoDetailsEdit(ctx context.Context, req *demo.DetailsEditReq) (res *demo.DetailsEditRes, err error) {
	err = service.DemoDetails().DemoDetailsEdit(ctx, req)
	return
}
func (c *demoDetailsController) DemoDetailsDetele(ctx context.Context, req *demo.DetailsDeleteReq) (res *demo.DetailsDeleteRes, err error) {
	err = service.DemoDetails().DemoDetailsDetele(ctx, req)
	return
}
func (c *demoDetailsController) DemoJiben(ctx context.Context, req *demo.JibenReq) (res *demo.JibenRes, err error) {
	res = new(demo.JibenRes)
	res.List, err = service.DemoDetails().DemoJiben(ctx, req.UserId)
	return
}
func (c *demoDetailsController) DemoGanbu(ctx context.Context, req *demo.GanbuReq) (res *demo.GanbuRes, err error) {
	res = new(demo.GanbuRes)
	res.List, err = service.DemoDetails().DemoGanbu(ctx, req)
	return
}
