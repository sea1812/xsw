package test

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"time"
	"xsw/app"
)

func TestPosts(r *ghttp.Request) {
	mPosts := app.TPosts{
		PageCount:     1,
		CurrentPage:   0,
		SQL:           "select * from posts",
		KeyField:      "",
		KeyValue:      "",
		CacheEnabled:  true,
		CacheKey:      "posts",
		CacheDuration: time.Second * 60,
	}
	mPosts.LoadPostsFromDB()
	fmt.Println(len(mPosts.Items))
	r.Response.Write(mPosts.Items)
}
