package main

import (
	"github.com/gogf/gf/frame/g"
	"xsw/app"
)

func main() {
	s := g.Server()

	//检查系统环境配置

	//读取全局变量

	//初始化路由
	s.AddStaticPath("/static", "./static")
	GroupFront := s.Group("/")
	GroupFront.ALL("/", app.PageIndex)
	s.SetPort(9001)
	s.Run()
}
