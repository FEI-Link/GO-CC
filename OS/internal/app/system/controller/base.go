/*
* @desc:system base controller

 */

package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

type BaseController struct {
	commonController.BaseController
}

// Init 自动执行的初始化方法
func (c *BaseController) Init(r *ghttp.Request) {
	c.BaseController.Init(r)
}