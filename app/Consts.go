package app

import "time"

/*
	常量和全局变量定义
*/

type TCacheVar struct {
	CacheKey string        //缓存名
	Duration time.Duration //缓存时间
}

var (
	//系统设置的缓存设置
	SysConfigCache = TCacheVar{
		CacheKey: "SystemConfig",
		Duration: time.Minute * 12,
	}
	//系统设置保存的文件名
	SysConfigFilename string = "./data/config.json"
)
