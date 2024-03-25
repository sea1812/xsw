package app

import (
	"github.com/gogf/gf/net/ghttp"
	_ "github.com/mattn/go-sqlite3"
)

func PageIndex(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/test/index.html")
}
