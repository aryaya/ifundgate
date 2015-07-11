//
//
//

package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/wangch/icloudfund/routers"
	"log"
)

func main() {
	beego.EnableHttpListen = false
	beego.EnableHttpTLS = true
	beego.HttpsPort = 443
	beego.HttpCertFile = "cert.pem"
	beego.HttpKeyFile = "key.pem"

	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/ripple.txt", "./ripple.txt")
	beego.SetStaticPath("/favicon.png", "./favicon.png")

	conf, err := readIniFile("ripple.txt")
	if err != nil {
		log.Fatal(err)
	}

	beego.Get("/federation", func(ctx *context.Context) {
		federation(ctx, conf)
	})

	// beego.Get("/quote", func(ctx *context.Context) {
	// 	quote(ctx, conf)
	// })

	beego.Run()
}
