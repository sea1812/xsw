package test

import (
	"github.com/gogf/gf/net/ghttp"
	"xsw/app"
)

func TestConfigV2(r *ghttp.Request) {
	m := app.TConfigV2{
		Logo:            "/static/logo.svg",
		BaseDomain:      "ggyun.top",
		BasePort:        "80",
		SiteTitle:       "技术笔记",
		SiteDescription: "我的技术笔记",
		SiteKeywords:    "golang,rust,react",
		Theme:           "default",
		ThemePath:       "./template/default",
		RunAsStatic:     false,
		UseRedis:        false,
		CacheEnable:     false,
		ICPSerial:       "京ICP备111111",
		Copyright:       "Copyright by ggyun.top 2024",
		NavShowArchive:  false,
		NavShowAbout:    false,
	}
	m.SaveToFile("./data/config.json")
	m2 := app.TConfigV2{}
	m2.LoadFromFile("./data/config.json")
	_ = r.Response.WriteJson(m2)
}
