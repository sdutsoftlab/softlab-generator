package gen

import (
	"github.com/sdutsoftlab/softlab-generator/utils"
	"os"
)

// 网站配置文件
var conf Website

func Init() {
	log.Info("开始初始化完成")
	// 初始化生成目录，把assets文件夹从主题文件夹中复制过去
	distInit()
	// 初始化md目录
	postsInit()
	log.Info("站点初始化完成")
}

func ConfInit() {
	conf = Config()
	//fmt.Println(conf)
}

func distInit() {
	err := os.MkdirAll(conf.Dist+"assets/", os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = utils.CopyDir(conf.Theme+"assets/", conf.Dist+"assets/")
	if err != nil {
		panic(err)
	}
}

func postsInit() {
	err := os.MkdirAll(conf.Markdown, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
