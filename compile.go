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
author:
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
	compileArticle()
	compileHome()
	log.Debug("编译完成")
}

func compileHome() {
	data["artlist"] = GetHomeArt()

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

func compileArticle() {
	LoadArticle()
	for _, post := range articles {

		data["title"] = post.Title

		data["article"] = post

		filepath := path.Join(conf.Dist, post.Url)
		fmt.Println(filepath)
		err := utils.Mkdir(filepath)
		if err != nil {
			panic(err)
		}

		file := path.Join(filepath, "index.html")

		htmlFile, err := os.Create(file)
		if err != nil {
			panic(err)
		}
		t, err := template.New("main.tpl").Funcs(funcMap).ParseFiles(
			conf.Theme+"layout/article.tpl",
			conf.Theme+"layout/main.tpl")
		if err != nil {
			panic(err)
		}

		err = t.Execute(htmlFile, data)
		if err != nil {
			panic(err)
		}

	}
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
