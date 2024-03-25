package app

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/os/gfile"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

/*
	前端页面类
*/

type TFrontPage struct {
	Config           TConfig //默认设置
	TemplateFilename string  //模板名称
	CustomData       g.Map   //用户传入的数据
}

// Display 显示页面，输出到浏览器
func (p *TFrontPage) Display(r *ghttp.Request) {
	mTemplateFilename := gfile.Join(p.Config.ItemByKey("site_theme").Value, p.TemplateFilename) //组合最终的模板文件名称
	if gfile.Exists(mTemplateFilename) {
		//模板文件存在，正常输出模板
		_ = r.Response.WriteTpl(mTemplateFilename, g.Map{
			"Config": p.Config,
			"Data":   p.CustomData,
		})
	} else {
		//模板文件不存在，则输出404页面丢失页面
		m404TemplateFilename := gfile.Join(p.Config.ItemByKey("site_theme").Value, "404.html")
		if gfile.Exists(m404TemplateFilename) == false {
			//如果主题中不存在404模板，则使用系统的404模板进行输出
			m404TemplateFilename = "404.html"
		}
		_ = r.Response.WriteTpl(m404TemplateFilename, g.Map{
			"Config": p.Config,
			"Data":   p.CustomData,
		})
	}
}

// RenderToFile 渲染页面，输出到文件
func (p *TFrontPage) RenderToFile(AFilename string) {
	mTemplateFilename := gfile.Join(p.Config.ItemByKey("site_theme").Value, p.TemplateFilename) //组合最终的模板文件名称
	var mC string
	if gfile.Exists(mTemplateFilename) {
		//模板文件存在，正常输出模板
		mC, _ = g.View().Parse(gctx.New(), mTemplateFilename, g.Map{
			"Config": p.Config,
			"Data":   p.CustomData,
		})
	} else {
		//模板文件不存在，则输出404页面丢失页面
		m404TemplateFilename := gfile.Join(p.Config.ItemByKey("site_theme").Value, "404.html")
		if gfile.Exists(m404TemplateFilename) == false {
			//如果主题中不存在404模板，则使用系统的404模板进行输出
			m404TemplateFilename = "404.html"
		}
		mC, _ = g.View().Parse(gctx.New(), m404TemplateFilename, g.Map{
			"Config": p.Config,
			"Data":   p.CustomData,
		})
	}
	//保存文件
	if strings.TrimSpace(mC) != "" {
		_ = gfile.PutContents(AFilename, mC)
	}
}

func PageIndex(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/test/index.html")
}
