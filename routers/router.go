package routers

import (
	"github.com/astaxie/beego"
	"github.com/wangch/icloudfund/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/deposit", &controllers.MainController{}, "get:Deposit")
	beego.Router("/deposit/*", &controllers.MainController{}, "post:Deposit")
}
