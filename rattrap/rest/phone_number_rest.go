package rest

import (
	"github.com/kataras/iris/v12"
	"mouse/rattrap/common"
	"mouse/rattrap/request"
)

//
// 检测服务
//

func PhoneDetection(ctx iris.Context) {
	var numbers request.PhoneDetectionRequest
	_ = ctx.ReadJSON(&numbers)
	ctx.JSON(common.Success(numbers))
}
