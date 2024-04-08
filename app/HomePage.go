package app

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"time"
)

/*
	HomePage类
*/

type THomepage struct {
	TFrontPageV2        //继承结构
	Posts        TPosts //帖子列表
}

// Init 初始化首页
func (p *THomepage) Init(r *ghttp.Request) {
	//设置首页属性
	p.PageTitle = "技术笔记"
	p.CustomHead = ""
	p.BaseTemplate = gfile.Join("./", p.Config.Theme) + "/index.html"
	p.ContentTemplate = gfile.Join(gfile.Join("./", p.Config.Theme) + "/homepage.html")
	p.AttachTemplate = ""
	p.CacheEnable = true
	p.CacheKey = "hompage"
	p.CacheDuration = time.Minute * 60
	p.Request = r

	//加载首页帖子
	/*
		p.Posts.SQL = "select * from posts order by modifytime desc"
		p.CacheEnable = true
		p.CacheKey = "homepage_posts"
		p.CacheDuration = time.Second * 60
		p.Posts.LoadPostsFromDB()
	*/
}

// PageHomepage 对接URL路由的函数
func PageHomepage(r *ghttp.Request) {
	mHomepage := THomepage{}
	mHomepage.Config.LoadFromFile("./data/config.json")
	mHomepage.Init(r)
	mHomepage.CacheEnable = false
	//获取当前页码
	mCurrentPage := r.GetInt("p")
	//获取全部页数
	mTotalPages := 1
	//刷新首页帖子列表
	mHomepage.Posts.SQL = fmt.Sprintf("select * from posts order by modifytime desc limit %d offset %d", 10, 10*mCurrentPage)
	mHomepage.Posts.CurrentPage = mCurrentPage
	mHomepage.Posts.PageCount = mTotalPages
	mHomepage.Posts.CacheEnabled = true
	mHomepage.Posts.CacheKey = "homepage_posts_" + fmt.Sprint(mCurrentPage)
	mHomepage.Posts.CacheDuration = time.Second * 60
	mHomepage.Posts.LoadPostsFromDB()
	mHomepage.CustomData = g.Map{
		"posts": mHomepage.Posts,
	}
	mHomepage.Display()
	fmt.Println(mHomepage.outData)
}
