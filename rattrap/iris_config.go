package rattrap

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"log"
	"mouse/rattrap/config"
	"mouse/rattrap/rest"
)

//
// Iris的配置启动
//

func init() {
	//start()
}

func start() {
	log.Println("Init Iris web server!")
	app := iris.New()
	app = route(app)
	err := app.Run(iris.Addr(config.Conf.Server.Host), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		log.Fatalf("Server start error failed %s", err)
	}
}

func route(app *iris.Application) *iris.Application {
	log.Println("Add route path to Iris web server!")
	party := app.Party("/v1")
	party.PartyFunc("/phone/phone-detection", func(phoneParty router.Party) {
		phoneParty.Post("/detection", rest.PhoneDetection)
	})
	return app
}
