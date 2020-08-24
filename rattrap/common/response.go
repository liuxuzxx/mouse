package common

import "github.com/kataras/iris/v12"

func Success(data interface{}) iris.Map {
	return iris.Map{"code": 1, "message": "success", "data": data}
}
