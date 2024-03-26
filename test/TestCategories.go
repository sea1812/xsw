package test

import (
	"github.com/gogf/gf/net/ghttp"
	"xsw/app"
)

func TestCategories(r *ghttp.Request) {
	mC := app.TCategories{}
	mC.Table = "categories"
	mC.LoadFromDB()
	r.Response.Write(mC.Items)
}
