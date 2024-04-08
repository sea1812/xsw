package app

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"time"
)

/*
	前端页面类第2版
*/

type TFrontPageV2 struct {
	Config          TConfigV2      //系统设置
	State           TPageState     //页面状态
	PageTitle       string         //页面标题
	CustomHead      string         //自定义Head内容
	BaseTemplate    string         //基础模板文件名
	ContentTemplate string         //内容模板文件名
	AttachTemplate  string         //附加模板文件名
	CacheEnable     bool           //是否缓存页面
	CacheKey        string         //缓存页面的KEY名
	CacheDuration   time.Duration  //缓存页面的时间
	startTime       time.Time      //开始生成页面的时间
	endTime         time.Time      //完成生成页面的时间
	CustomData      g.Map          //用户传入的数据
	Request         *ghttp.Request //传入的请求对象
	outData         g.Map          //传输到模板系统的MAP
}

// Init 初始化
func (p *TFrontPageV2) Init() {
	//赋值开始生成页面时间
	p.startTime = time.Now()
	//加载系统设置信息
	p.Config.LoadFromFile(SysConfigFilename)
	if p.CustomData == nil {
		p.CustomData = make(g.Map)
	}
	p.outData = make(g.Map)
}

// PrepareOutdata 准备用于输出的OutData
func (p *TFrontPageV2) PrepareOutdata() {
	if p.outData == nil {
		p.Init()
	}
	p.outData["System"] = p.Config.ToMap()
	p.outData["State"] = p.State
	p.outData["PageTitle"] = p.PageTitle
	p.outData["ContentTpl"] = p.ContentTemplate
	p.outData["AttachTpl"] = p.AttachTemplate
	p.outData["CustomHead"] = p.CustomHead
	p.outData["Data"] = p.CustomData
	p.outData["State"] = p.State
}

// RenderPage 生成页面
func (p *TFrontPageV2) RenderPage() string {
	p.PrepareOutdata() //准备用于输出的OutData
	var mR string = ""
	mR, _ = p.Request.Response.ParseTpl(p.BaseTemplate, p.outData)
	fmt.Println("------------", p.outData)
	return mR
}

// Display 显示页面
func (p *TFrontPageV2) Display() {
	var mPage string //页面内容
	p.PrepareOutdata()
	if p.CacheEnable {
		//使用页面缓存
		mIs, _ := gcache.Contains(p.CacheKey)
		if mIs == true {
			//缓存存在，从缓存中读取
			m2, _ := gcache.Get(p.CacheKey)
			mPage = fmt.Sprint(m2)
		} else {
			//缓存不存在，生成页面并存入缓存
			mPage = p.RenderPage()
			p.endTime = time.Now()
			p.outData["SubTime"] = p.endTime.Sub(p.startTime)
			_ = gcache.Set(p.CacheKey, mPage, p.CacheDuration)
		}
	} else {
		//不使用页面缓存，直接生成页面
		p.endTime = time.Now()
		p.outData["SubTime"] = p.endTime.Sub(p.startTime)
		mPage = p.RenderPage()
	}
	//输出页面内容
	p.Request.Response.Write(mPage)
}
