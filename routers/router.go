package routers

import (
	"github.com/astaxie/beego"
	"github.com/wangch/icloudfund/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
