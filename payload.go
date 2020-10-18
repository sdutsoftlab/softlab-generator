package gen

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/sdutsoftlab/softlab-generator/utils"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	articles []*Article
)

// 加载所有的markdown文件
func LoadArticle() {
	articles = make([]*Article, 0)
	mdList := MDList() // 路径名列表
	for _, md := range mdList {
		post, err := loadContent(md)

		if err == nil {
			post.Url = createPostLink(post)
			articles = append(articles, post)
		} else {
			panic(err)
		}
	}
	sort.Sort(Articles(articles))
	//for i, v := range articles {
	//	fmt.Println(i, v)
	//}
}

// 获取首页文章
func GetHomeArt() []*Article {
	homeArt := make([]*Article, len(articles))
	copy(homeArt, articles)
	return homeArt
}

// 获取 markdown 文件夹下所有.md文件的路径名
func MDList() []string {
	posts := path.Join(conf.Dist, "posts/")
	err := utils.Mkdir(posts)
	if err != nil {
		panic(err)
	}

	mds := make([]string, 0)
	err = filepath.Walk(conf.Markdown, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".md") {
			mds = append(mds, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return mds
}

type mustring struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Date        string   `yaml:"date"`
	Time        string   `yaml:"time"`
	Author      string   `yaml:"author"`
	Categories  []string `yaml:"categories"`
	Tags        []string `yaml:"tags"`
	Content     string   `yaml:"content"`
}

// 加载markdown内容，并转成html
func loadContent(file string) (*Article, error) {
	ctx, err := readMuCtx(file)
	if err != nil {
		return nil, err
	}
	a := &Article{
		Title:     ctx.Title,
		Author:    ctx.Author,
		Content:   utils.MarkdownToHtml(ctx.Content),
		Tags:      ctx.Tags,
		Category:  ctx.Categories,
		CreatedAt: utils.Str2Unix("2006-01-02", ctx.Date),
	}
	return a, nil
}

// 创建连接
func createPostLink(post *Article) string {
	t := time.Unix(post.CreatedAt, 0)
	year, month, day := t.Date()
	return fmt.Sprintf("/%s/%d/%d/%d/%s/", "posts", year,
		month, day, utils.Convert(post.Title))
}

// 解析读取md开头配置 和 内容
func readMuCtx(file string) (ctx *mustring, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	firstLine, err := reader.ReadString('\n') // 读取到换行结束
	if err != nil {
		return nil, err
	}

	// 第一行不是以 ---起首error
	if !strings.HasPrefix(firstLine, "---") {
		err = fmt.Errorf("markdown file format error, the file header must start with '---' : " + file)
		return nil, err
	}

	// 读取---   到   ---之间的内容
	buf := bytes.NewBuffer(nil)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
		}
		if strings.HasPrefix(line, "---") {
			break
		} else {
			if err == io.EOF {
				err = fmt.Errorf("markdown file format error, the file header must end with '---' : " + file)
				return nil, err
			}
		}
		buf.WriteString(line)
	}

	// 反序列化头部配置
	err = yaml.Unmarshal(buf.Bytes(), &ctx)
	if err != nil {
		return nil, err
	}

	//全部读取未读取部分
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	stat, _ := f.Stat()

	//标题为空，默认使用文件名字作为文章标题
	if ctx.Title == "" {
		ctx.Title = strings.TrimRight(stat.Name(), ".md")
	}
	//时间为空，默认使用文章的修改时间
	if ctx.Date == "" {
		ctx.Date = utils.Format(stat.ModTime().Unix())
	}

	ctx.Content = string(content)
	return
}
