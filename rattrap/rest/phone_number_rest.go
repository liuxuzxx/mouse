package rest

import (
	"github.com/kataras/iris/v12"
	"mouse/rattrap/common"
	"mouse/rattrap/redis"
	"mouse/rattrap/request"
	"mouse/rattrap/service"
	"strconv"
	"time"
)

//
// 检测服务
//

func PhoneDetection(ctx iris.Context) {
	second := time.Now().Unix()
	redis.Incrby(strconv.FormatInt(second,10), 1)
	var number request.PhoneDetectionRequest
	_ = ctx.ReadJSON(&number)
	_, _ = ctx.JSON(common.Success(service.PhoneDetectionServiceImpl.Detection(number)))
}
