package test

import (
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
		CacheEnable:     true,
		CacheKey:        "testpage",
		CacheDuration:   time.Minute * 60,
		CustomData:      nil,
		Request:         r,
	}
	mPage.Init()
	mPage.Display()
}
