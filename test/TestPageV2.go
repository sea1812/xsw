package test

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
	"xsw/app"
)

func TestPageV2(r *ghttp.Request) {
	mPage := app.TFrontPageV2{
		PageTitle:       "技术笔记",
		CustomHead:      "",
		BaseTemplate:    "/test/index.html",
		ContentTemplate: "",
		AttachTemplate:  "",
		CacheEnable:     false,
		CacheKey:        "testpage",
		CacheDuration:   time.Minute * 60,
		CustomData: g.Map{
			"Demo": "This is demo",
		},
		Request: r,
	}
	mPage.Init()
	mPage.Display()
}
