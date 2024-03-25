package test

import (
	"fmt"
	"xsw/app"
)

func TestClass() {
	var mConfig app.TConfig
	//设置表名
	mConfig.Table = "config"
	//从数据库加载
	mConfig.LoadFromDb()
	//打印所有内容
	for _, v := range mConfig.Data {
		fmt.Println(v)
	}
}
