package app

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
)

/*
	系统设置对象V2版本
*/

type TConfigV2 struct {
	Logo            string `json:"logo"`            //Logo图片路径和名称
	BaseDomain      string `json:"baseDomain"`      //基准域名
	BasePort        string `json:"basePort"`        //基准端口
	SiteTitle       string `json:"siteTitle"`       //网站标题
	SiteDescription string `json:"siteDescription"` //网站描述（SEO用）
	SiteKeywords    string `json:"siteKeywords"`    //网站搜索关键字（SEO用）
	Theme           string `json:"defaultTheme"`    //主题名称
	ThemePath       string `json:"themePath"`       //主题路径
	RunAsStatic     bool   `json:"runAsStatic"`     //是否以静态网站模式运行
	UseRedis        bool   `json:"useRedis"`        //是否使用Redis缓存
	CacheEnable     bool   `json:"cacheEnable"`     //是否使用缓存
	ICPSerial       string `json:"ICPSerial"`       //ICP备案号
	Copyright       string `json:"copyright"`       //版权声明信息字符串
	NavShowArchive  bool   `json:"navShowArchive"`  //导航栏是否显示归档菜单
	NavShowAbout    bool   `json:"navShowAbout"`    //导航栏是否显示关于菜单
}

// SaveToFile 保存到文件
func (p *TConfigV2) SaveToFile(AFilename string) {
	mJson := gjson.New(p)
	_ = gfile.PutContents(AFilename, mJson.Export())
	if p.CacheEnable == true {
		//写入Cache
		_ = gcache.Set(SysConfigCache.CacheKey, mJson.Export(), SysConfigCache.Duration)
	}
}

// LoadFromFile 从文件中读取
func (p *TConfigV2) LoadFromFile(AFilename string) {
	mIs, _ := gcache.Contains(SysConfigCache.CacheKey)
	if mIs == true {
		mJsonStr, _ := gcache.Get(SysConfigCache.CacheKey)
		mJson := gjson.New(mJsonStr)
		_ = mJson.UnmarshalValue(p)
	} else {
		mJson, er := gjson.Load(AFilename, true)
		fmt.Println(mJson, er)
		_ = mJson.UnmarshalValue(p)
	}
}
