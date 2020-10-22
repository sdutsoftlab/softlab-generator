package utils

import (
	//bf "github.com/russross/blackfriday"
	bf "gopkg.in/russross/blackfriday.v2"
	"log"
	"regexp"
	"strings"
)

// 封装Markdown转换为Html的逻辑
var (
	TOC_TITLE = "<h4>目录:</h4>"
)

var navRegex = regexp.MustCompile(`(?ismU)<nav>(.*)</nav>`)

func MarkdownToHtml(content string) (str string) {
	defer func() {
		e := recover()
		if e != nil {
			str = content
			log.Println("Render Markdown ERR:", e)
		}
	}()

	htmlFlags := bf.HTMLFlags(0)

	if strings.Contains(strings.ToLower(content), "[toc]") {

		htmlFlags |= bf.TOC
	}

	htmlFlags |= bf.UseXHTML
	htmlFlags |= bf.Smartypants
	htmlFlags |= bf.SmartypantsFractions
	htmlFlags |= bf.SmartypantsLatexDashes
	htmlFlags |= bf.FootnoteReturnLinks

	renderer := bf.NewHTMLRenderer(bf.HTMLRendererParameters{Flags: htmlFlags})

	extensions := bf.Extensions(0)
	extensions |= bf.NoIntraEmphasis
	extensions |= bf.Tables
	extensions |= bf.FencedCode
	extensions |= bf.Autolink
	extensions |= bf.Strikethrough
	extensions |= bf.SpaceHeadings
	extensions |= bf.HardLineBreak
	extensions |= bf.Footnotes

	str = string(bf.Run([]byte(content), bf.WithRenderer(renderer), bf.WithExtensions(extensions)))

	if htmlFlags&bf.TOC != 0 {
		found := navRegex.FindIndex([]byte(str))
		if len(found) > 0 {
			toc := str[found[0]:found[1]]
			toc = TOC_TITLE + toc
			str = str[found[1]:]
			reg := regexp.MustCompile(`\[toc\]|\[TOC\]`)
			str = reg.ReplaceAllString(str, toc)
		}
	}
	return str
}
