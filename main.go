package main

import (
	"github.com/gogf/gf/frame/g"
	"xsw/app"
	"xsw/test"
)

func main() {
	s := g.Server()
	s.SetDumpRouterMap(false)
	//检查系统环境配置
	test.TestClass()
	//读取全局变量

	//初始化路由
	s.AddStaticPath("/static", "./static")
	GroupFront := s.Group("/")
	GroupFront.ALL("/", app.PageIndex)

	//测试页面
	GroupTest := s.Group("/test")
	GroupTest.ALL("/", test.TestingPage)

	//s.SetPort(9001)
	s.Run()
}
