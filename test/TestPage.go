package test

import (
	"github.com/gogf/gf/net/ghttp"
	"xsw/app"
)

// TestingPage TestPage 测试TFrontPage
func TestingPage(r *ghttp.Request) {
	mPage := app.TFrontPage{}
	mPage.Init()
	mPage.TemplateFilename = "index.html"
	mPage.Display(r)
}
