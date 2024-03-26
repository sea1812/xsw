package app

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

/*
	管理和记录页面的状态，状态记载到session中
*/

type TPageState struct {
	Request  *ghttp.Request
	PageId   string `json:"pageId"`   //页面ID
	IsAdmin  bool   `json:"isAdmin"`  //当前是否admin
	IsMember bool   `json:"isMember"` //当前是否普通会员
}

// Load 刷新数据，即从session中重载状态信号
func (p *TPageState) Load() {
	if p.Request.Session.Contains(p.PageId) {
		//取存储的字符串
		mJS := p.Request.Session.GetString(p.PageId)
		//转换成JSON
		mJ := gjson.New(mJS)
		p.IsAdmin = mJ.GetBool("isAdmin")
		p.IsMember = mJ.GetBool("isMember")
	} else {
		p.IsAdmin = false
		p.IsMember = false
	}
}

// SetState 设置状态数据
func (p *TPageState) SetState(AIsAdmin bool, AIsMember bool) {
	p.IsMember = AIsMember
	p.IsAdmin = AIsAdmin
}

// Save 保存状态数据到session
func (p *TPageState) Save() {
	//组合生成JSON字符串
	mJ := gjson.New(g.Map{
		"isAdmin":  p.IsAdmin,
		"isMember": p.IsMember,
	})
	_ = p.Request.Session.Set(p.PageId, mJ.Export())
}
