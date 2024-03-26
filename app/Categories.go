package app

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"strings"
	"time"
)

/*
	分类目录对象
*/

type TCategoryItem struct {
	Id          int    `json:"id"`          //ID
	Caption     string `json:"caption"`     //标题
	Position    int    `json:"position"`    //排序
	Url         string `json:"url"`         //链接
	Description string `json:"description"` //简介
	Banner      string `json:"banner"`      //标题图片
	IsHidden    int    `json:"isHidden"`    //是否隐藏分类，隐藏分类只有站长登陆才能查看
	IsRegister  int    `json:"isRegister"`  //是否必须注册用户登陆后才能查看
	IsSecret    int    `json:"is_secret"`   //是否密码保护频道，必须输入密码才能查看
	SecretPwd   string `json:"secretPwd"`   //密码保护频道的访问密码
}

type TCategories struct {
	Table string           //数据库表名
	Items []*TCategoryItem //所有分类的列表
}

// SaveToDB 保存分类表到数据库
func (p *TCategories) SaveToDB() {
	if strings.TrimSpace(p.Table) != "" {
		for _, v := range p.Items {
			if v.Id == -1 {
				//插入
				_, _ = g.DB().Model("categories").Insert(g.Map{
					"caption":     v.Caption,
					"position":    v.Position,
					"url":         v.Url,
					"banner":      v.Banner,
					"description": v.Description,
					"is_hidden":   v.IsHidden,
					"is_secret":   v.IsSecret,
					"is_register": v.IsRegister,
					"secret_pwd":  v.SecretPwd,
				})
			} else {
				//按ID更新
				_, _ = g.DB().Model("categories").Update(g.Map{
					"caption":     v.Caption,
					"position":    v.Position,
					"url":         v.Url,
					"banner":      v.Banner,
					"description": v.Description,
					"is_hidden":   v.IsHidden,
					"is_secret":   v.IsSecret,
					"is_register": v.IsRegister,
					"secret_pwd":  v.SecretPwd,
				}, "id=?", v.Id)
			}
		}
	}
}

// LoadFromDB 从数据库加载分类目录
func (p *TCategories) LoadFromDB() {
	p.Clear()
	res, _ := g.DB().Model("categories").OrderDesc("position").Cache(time.Second*30, "categories").All()
	for _, v := range res {
		fmt.Println(v)
		mItem := p.Add()
		mItem.Url = v["url"].String()
		mItem.Id = v["id"].Int()
		mItem.Position = v["position"].Int()
		mItem.Banner = v["banner"].String()
		mItem.Description = v["description"].String()
		mItem.IsRegister = v["is_register"].Int()
		mItem.IsHidden = v["is_hidden"].Int()
		mItem.IsSecret = v["is_secret"].Int()
		mItem.SecretPwd = v["secret_pwd"].String()
		fmt.Println(mItem)
	}
}

// Clear 清除全部分类表
func (p *TCategories) Clear() {
	p.Items = nil
}

// Add 增加分类
func (p *TCategories) Add() *TCategoryItem {
	mItem := TCategoryItem{}
	p.Items = append(p.Items, &mItem)
	return &mItem
}

// ItemById 按ID查找分类项
func (p *TCategories) ItemById(AID int) *TCategoryItem {
	var mR *TCategoryItem = nil
	for _, v := range p.Items {
		if v.Id == AID {
			mR = v
			break
		}
	}
	return mR
}

// ItemByCaption 按Caption名查找分类项
func (p *TCategories) ItemByCaption(ACaption string) *TCategoryItem {
	var mR *TCategoryItem = nil
	for _, v := range p.Items {
		if v.Caption == ACaption {
			mR = v
			break
		}
	}
	return mR
}

// Init 初始化
func (p *TCategories) Init() {
	if strings.TrimSpace(p.Table) != "" {
		p.LoadFromDB()
	}
}
