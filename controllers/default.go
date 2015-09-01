//
//
//

package controllers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/wangch/glog"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "index.html"
}

func (c *MainController) DepositAmount() {
}

func (c *MainController) Deposit() {
	r := c.Ctx.Request
	u := *r.URL
	u.Scheme = "http"
	u.Host = Gconf.Host
	u.Path = "/api" + u.Path

	glog.Infoln(u)

	var resp *http.Response
	var err error
	if r.Method == "GET" {
		resp, err = http.Get(u.String())
	} else {
		contentType := r.Header.Get("Content-Type")
		var nr *http.Request
		if strings.Contains(contentType, "multipart/form-data") {
			buf := &bytes.Buffer{}
			writer := multipart.NewWriter(buf)
			i := strings.Index(contentType, "--")
			boundary := contentType[i:]
			err = writer.SetBoundary(boundary)
			if err != nil {
				glog.Infoln(err)
				return
			}

			for k, v := range r.MultipartForm.Value {
				writer.WriteField(k, v[0])
			}
			for k, v := range r.MultipartForm.File {
				w, err := writer.CreateFormFile(k, v[0].Filename)
				if err != nil {
					glog.Errorln(err)
					return
				}
				f, err := v[0].Open()
				if err != nil {
					glog.Errorln(err)
					return
				}
				io.Copy(w, f)
				f.Close()
			}
			writer.Close()
			nr, err = http.NewRequest(r.Method, u.String(), buf)
		} else {
			nr, err = http.NewRequest(r.Method, u.String(), strings.NewReader(r.Form.Encode()))
		}
		if err != nil {
			glog.Errorln(err)
			return
		}
		nr.Header.Set("Content-Type", contentType)
		resp, err = http.DefaultClient.Do(nr)
	}

	if err != nil {
		glog.Errorln(err)
		c.Redirect("/", 302)
		return
	}
	io.Copy(c.Ctx.ResponseWriter, resp.Body)
}
