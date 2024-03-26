package app

import (
	"github.com/gogf/gf/frame/g"
	"strings"
	"time"
)

// TConfigItem 设置项目对象
type TConfigItem struct {
	Key     string //Key名称
	Value   string //Value值
	Comment string //备注
}

// TConfig 设置对象
type TConfig struct {
	Table string        //持久化对应的SQLite表
	Data  []TConfigItem //保存所有设置项目对
}

// Add 增加设置项目，返回对象指针
func (p *TConfig) Add(AKey string, AValue string, AComment string) *TConfigItem {
	mKey := TConfigItem{
		Key:     AKey,
		Value:   AValue,
		Comment: AComment,
	}
	p.Data = append(p.Data, mKey)
	return &mKey
}

// ItemByKey 根据Key名查找设置项
func (p *TConfig) ItemByKey(AKey string) *TConfigItem {
	var mR *TConfigItem = nil
	for _, v := range p.Data {
		if v.Key == AKey {
			mR = &v
			break
		}
	}
	return mR
}

// Clear 清除全部设置项目
func (p *TConfig) Clear() {
	p.Data = nil
}

// 检查数据库中是否存在指定Key的记录
func (p *TConfig) ifKeyExists(AKey string) bool {
	mR := false
	if strings.TrimSpace(p.Table) != "" {
		res, _ := g.DB().GetCount("select * from "+p.Table+" where key=?", AKey)
		mR = res != 0
	}
	return mR
}

// SaveToDb 保存到数据库
func (p *TConfig) SaveToDb() {
	if strings.TrimSpace(p.Table) != "" {
		for _, v := range p.Data {
			if p.ifKeyExists(v.Key) == false {
				//插入
				_, _ = g.DB().Model(p.Table).Insert(g.Map{
					"key":     v.Key,
					"value":   v.Value,
					"comment": v.Comment,
				})
			} else {
				//更新
				_, _ = g.DB().Model(p.Table).Update(g.Map{
					"value":   v.Value,
					"comment": v.Comment,
				}, "where key=?", v.Key)
			}
		}
	}
}

// LoadFromDb 从数据库中读取
func (p *TConfig) LoadFromDb() {
	if strings.TrimSpace(p.Table) != "" {
		//数据库设置5分钟缓存
		res, _ := g.DB().Model(p.Table).Cache(time.Minute*5, "CONFIG_"+p.Table).All()
		p.Clear()
		for _, v := range res {
			p.Add(v["key"].String(), v["value"].String(), v["comments"].String())
		}
	}
}

// ToMap Map 导出成Map
func (p *TConfig) ToMap() *g.Map {
	var mR g.Map = make(g.Map)
	for _, v := range p.Data {
		mR[v.Key] = v.Value
	}
	return &mR
}

// GetValue 获取Key的值
func (p *TConfig) GetValue(AKey string) string {
	mItem := p.ItemByKey(AKey)
	if mItem != nil {
		return mItem.Value
	} else {
		return ""
	}
}

// SetValue 更新Key的值
func (p *TConfig) SetValue(AKey string, AValue string, AComment string) {
	mItem := p.ItemByKey(AKey)
	if mItem != nil {
		mItem.Value = AValue
		if strings.TrimSpace(AComment) != "" {
			mItem.Comment = AComment
		}
	}
}
