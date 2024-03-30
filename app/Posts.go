package app

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"time"
)

/*
	帖子管理类对象
*/

// TPostItem 帖子项对象
type TPostItem struct {
	Url           string        `json:"url"`           //帖子链接地址
	Id            int           `json:"id"`            //帖子ID
	Title         string        `json:"title"`         //标题
	Banner        string        `json:"banner"`        //Banner图片路径
	Category      string        `json:"category"`      //所属分类名称
	CategoryId    int           `json:"categoryId"`    //所属分类ID
	PublishTime   time.Time     `json:"publishTime"`   //发布时间
	ModifyTime    time.Time     `json:"modifyTime"`    //修改时间
	Tags          string        `json:"tags"`          //自定义标签，用逗号分割
	ShortDesc     string        `json:"shortDesc"`     //内容摘要文本
	Content       string        `json:"content"`       //内容文本
	IsPublished   bool          `json:"isPublished"`   //是否已发布，否则为草稿
	IsPublic      bool          `json:"isPublic"`      //是否公开，否则为私有
	IsSecret      bool          `json:"isSecret"`      //是否密码保护，是则需要输入密码
	SecretCode    string        `json:"secretCode"`    //保护密码
	IsRegistered  bool          `json:"isRegistered"`  //是否需要会员资格，是则只对登陆会员可见
	IsHighlight   bool          `json:"isHighlight"`   //是否置顶
	AllowPull     bool          `json:"allowPull"`     //是否允许点赞
	AllowComment  bool          `json:"allowComment"`  //是否允许评论
	MarkdownFile  string        `json:"markdownFile"`  //Markdown原文路径
	RefrenceUrl   string        `json:"refrenceUrl"`   //转载原文链接
	CacheEnabled  bool          `json:"cacheEnabled"`  //是否缓存
	CacheKey      string        `json:"cacheKey"`      //缓存KEY
	CacheDuration time.Duration `json:"cacheDuration"` //缓存时间
}

// TPosts  帖子列表
type TPosts struct {
	Items         []TPostItem   //帖子列表数组
	PageCount     int           //总页数
	CurrentPage   int           //当前页码
	SQL           string        //执行的SQL语句
	KeyField      string        //高亮词所属字段（如类别、TAG等）
	KeyValue      string        //高亮词内容
	CacheEnabled  bool          //是否启用缓存
	CacheKey      string        //缓存名
	CacheDuration time.Duration //缓存时间
}

// Clear 清除所有帖子
func (p *TPosts) Clear() {
	p.Items = nil
}

// AddItemFromRecord AddRecord 将gdb.record加入Items
func (p *TPosts) AddItemFromRecord(ARecord gdb.Record) {
	var mItem TPostItem
	mItem.Id = ARecord["id"].Int()
	mItem.Url = ARecord["url"].String()
	mItem.Title = ARecord["title"].String()
	mItem.Banner = ARecord["banner"].String()
	mItem.Category = ARecord["category"].String()
	mItem.CategoryId = ARecord["categoryid"].Int()
	mItem.PublishTime = ARecord["publishtime"].Time()
	mItem.ModifyTime = ARecord["modifytime"].Time()
	mItem.Tags = ARecord["tags"].String()
	mItem.ShortDesc = ARecord["shortdesc"].String()
	mItem.IsPublished = ARecord["ispublished"].Bool()
	mItem.IsPublic = ARecord["ispublic"].Bool()
	mItem.IsSecret = ARecord["issecret"].Bool()
	mItem.IsRegistered = ARecord["isregistered"].Bool()
	mItem.IsHighlight = ARecord["ishighlight"].Bool()
	mItem.AllowPull = ARecord["allowpull"].Bool()
	mItem.AllowComment = ARecord["allowcomment"].Bool()
	p.Items = append(p.Items, mItem)
}

// LoadPostsFromDB 执行SQL，从数据库中读取帖子
func (p *TPosts) LoadPostsFromDB() {
	var res gdb.Result
	var er error
	//检查是否缓存
	if p.CacheEnabled == true {
		mIs, _ := gcache.Contains(p.CacheKey)
		if mIs == true {
			mTmp, _ := gcache.Get(p.CacheKey)
			res = mTmp.(gdb.Result)
		} else {
			res, er = g.DB().GetAll(p.SQL)
			_ = gcache.Set(p.CacheKey, res, p.CacheDuration)
		}
	} else {
		res, er = g.DB().GetAll(p.SQL)
	}
	if er == nil {
		//循环RES，加入Items
		for _, v := range res {
			p.AddItemFromRecord(v)
		}
	}
}
