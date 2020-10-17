package gen

import (
	"fmt"
)

// 网站配置文件
var conf Website

//var (
//	mdHead  = `---
//date: %s
//time: %s
//title: %s
//categories:
//-
//tags:
//-
//-
//---`
//)

func Init() {
	conf = Config()
	fmt.Printf("%+v\n", conf)
	//CreateMarkdown("122")
}
