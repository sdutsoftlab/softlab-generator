package gen

import (
	"fmt"
	"github.com/sdutsoftlab/softlab-generator/utils"
	"html/template"
	"os"
	"path"
	"strconv"
	"time"
)

var (
	mdHead = `---
title: %s
date: %s
time: %s
categories:
-
tags:
-
-
---`
	funcMap = template.FuncMap{
		"now":       utils.Now,
		"unescaped": utils.Unescaped,
		"format":    utils.Format,
	}
	data = map[string]interface{}{
		"title":       conf.Title,
		"subtitle":    conf.SubTitle,
		"description": conf.Description,
		"keywords":    conf.Keywords,
		"author":      conf.Author,
		"avatar":      conf.Avatar,
		"github":      conf.Github,
	}
)

func Compile() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic recoverd from: %v", r)
		}
	}()

	log.Info("开始编译博客")
	compileHome()
	log.Debug("编译完成")
}

func compileHome() {
	data["artlist"] = GetHomeArt()
	data["cate"] = "GetCate()"
	data["tags"] = "GetTag()"

	err := utils.Mkdir(conf.Dist)
	if err != nil {
		panic("生成目录创建错误")
	}

	index := path.Join(conf.Dist, "index.html")

	file, err := os.Create(index)
	if err != nil {
		panic(err)
	}

	t, err := template.New("main.tpl").Funcs(funcMap).ParseFiles(
		conf.Theme+"layout/main.tpl",
		conf.Theme+"layout/home.tpl")
	if err != nil {
		panic(err)
	}

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

func CompileAirtcle() {

}

func CreateMarkdown(filename string) string {
	year, month, _ := time.Now().Date()

	// **pwd/posts/year/month/123.md
	dir := path.Join(conf.Markdown, strconv.Itoa(year),
		strconv.Itoa(int(month)))

	msg, err := utils.CreateFile(dir, filename+".md")
	if err != nil {
		log.Fatal(msg, err.Error())
	}

	// 创建成功
	date := time.Now().Format("2006-01-02")
	now := time.Now().Format("15:04:05")
	// 默认以文件名字作为文章标题，可以随后修改
	mdHeadStr := fmt.Sprintf(mdHead, filename, date, now)

	err = utils.WriteFile(dir, filename+".md", mdHeadStr)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}

// 获取首页文章
func GetHomeArt() []Article {
	num := conf.HomeArtNum
	homeArt := make([]Article, 0)

	for i := 0; i < num; i++ {
		a := Article{
			Id:          i,
			Title:       strconv.Itoa(i),
			Description: "11",
			Summary:     "22",
			Content:     "333",
			Tags:        []string{"t1", "t2"},
			Category:    []string{"c1", "c2"},
			CreatedAt:   0,
			UpdatedAt:   0,
			Url:         "111",
		}
		homeArt = append(homeArt, a)
	}
	return homeArt
}
