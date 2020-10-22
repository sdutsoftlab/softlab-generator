package main

import (
	"flag"
	"fmt"
	gen "github.com/sdutsoftlab/softlab-generator"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var (
	log = logrus.WithFields(logrus.Fields{
		`MOD`: `gen`,
	})
)

const (
	HELP = `

Usage:

gen command [args...]

	初始化博客文件夹
    	gen init

	新建 markdown 文件
    	gen new filename

	编译博客
    	gen compile/c

    运行chca所有服务，包括内置服务器、监听器
    	gen run [port]

	`
)

func PrintUsage() {
	fmt.Println(HELP)
}

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) == 0 || len(args) > 3 {
		PrintUsage()
		os.Exit(1)
	}

	gen.ConfInit()
	switch args[0] {
	default:
		PrintUsage()
	case "init":
		gen.Init()
	case "new":
		if len(args) == 2 {
			name := args[1]
			markdown := gen.CreateMarkdown(name)
			log.Info("创建成功: ", markdown)
		} else {
			panic("缺少文件名")
		}
	case "compile", "c":
		if len(args) == 1 {
			gen.Compile()
		} else {
			PrintUsage()
			panic("not found")
		}
	case "run":
		gen.Compile()
		var port int = 9305
		if len(args) == 2 {
			p, err := strconv.Atoi(args[1])
			if err != nil {
				panic("port is wrong")
			}
			port = p
		}
		gen.ListenHTTPServer(port)
	}
}
