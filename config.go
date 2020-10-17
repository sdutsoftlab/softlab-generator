package gen

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

var (
	confile = "public/config.yml"
)

func load() ([]byte, error) {
	_, err := os.Stat(confile)
	if os.IsNotExist(err) {

	}

	file, err := os.Open(confile)
	if err != nil {

	}
	return ioutil.ReadAll(file)
}

func Config() Website {
	data, err := load()
	if err != nil {
		panic("加载配置文件失败，" + err.Error())
	}
	website := Website{}
	err = yaml.Unmarshal(data, &website)
	if err != nil {
		panic("解析配置文件失败，" + err.Error())
	}
	return website
}
