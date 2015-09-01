//
//
//

package main

import (
	"io"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/wangch/glog"
	"github.com/wangch/icloudfund/controllers"
	_ "github.com/wangch/icloudfund/routers"
)

func init() {
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
	beego.SessionOn = true

	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/favicon.png", "./favicon.png")
}

func main() {
	glog.SetLogDirs(".")
	glog.SetLogToStderr(true)
	conf := controllers.Gconf
	beego.Get("/federation", func(ctx *context.Context) {
		federation(ctx, conf)
	})

	beego.Get("/ripple.txt", func(ctx *context.Context) {
		f, err := os.Open("ripple.txt")
		if err != nil {
			glog.Fatal(err)
		}
		io.Copy(ctx.ResponseWriter, f)
		f.Close()
	})

	beego.Get("/quote", func(ctx *context.Context) {
		u := "http://" + conf.Host + "/api/quote?" + ctx.Request.URL.RawQuery
		glog.Infoln(u)
		r, err := http.Get(u)
		if err != nil {
			glog.Errorln(err)
			return
		}
		io.Copy(ctx.ResponseWriter, r.Body)
		r.Body.Close()
	})

	beego.Run()
}
