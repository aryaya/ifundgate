//
//
//

package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/wangch/icloudfund/routers"
	"io"
	"log"
	"os"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.EnableHttpListen = false
	beego.EnableHttpTLS = true
	beego.HttpsPort = 443
	beego.HttpCertFile = "cert.pem"
	beego.HttpKeyFile = "key.pem"

	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/fonts", "static/fonts")
	// beego.SetStaticPath("/ripple.txt", "./ripple.txt")
	beego.SetStaticPath("/favicon.png", "./favicon.png")

	conf, err := readIniFile("ripple.txt")
	if err != nil {
		log.Fatal(err)
	}

	beego.Get("/federation", func(ctx *context.Context) {
		federation(ctx, conf)
	})

	beego.Get("/ripple.txt", func(ctx *context.Context) {
		f, err := os.Open("ripple.txt")
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(ctx.ResponseWriter, f)
	})

	// beego.Get("/quote", func(ctx *context.Context) {
	// 	quote(ctx, conf)
	// })

	beego.Run()
}
