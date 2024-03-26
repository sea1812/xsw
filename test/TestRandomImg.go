package test

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func GetImgUrl() {
	r, _ := g.Client().Get("https://api.lyiqk.cn/scenery")
	fmt.Println("URL====", r.Request.URL, "HEADER====", r.Header)
}
