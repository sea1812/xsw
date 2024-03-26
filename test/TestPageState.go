package test

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"xsw/app"
)

func TestPageState(r *ghttp.Request) {
	mP := app.TPageState{
		Request:  r,
		PageId:   "test_page",
		IsMember: false,
		IsAdmin:  true,
	}

	mP.SetState(false, true)
	mP.Save()
	mP.Load()

	r.Response.Write(gconv.Map(mP))
}
