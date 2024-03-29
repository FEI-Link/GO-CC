/*
* @desc:公共接口相关
 */

package common

import "gocc/internal/app/common/model"

// PageReq 公共请求参数
type PageReq struct {
	model.PageReq
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
