package rest

import (
	"github.com/kataras/iris/v12"
	"mouse/rattrap/common"
	"mouse/rattrap/request"
	"mouse/rattrap/service"
)

//
// 检测服务
//

func PhoneDetection(ctx iris.Context) {
	var number request.PhoneDetectionRequest
	_ = ctx.ReadJSON(&number)
	_, _ = ctx.JSON(common.Success(service.PhoneDetectionServiceImpl.Detection(number)))
}
