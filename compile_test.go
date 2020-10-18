package gen

import "testing"

func TestCompileAirtcle(t *testing.T) {
	Init()
	compileArticle()
}

func TestCreateMarkdown(t *testing.T) {
	Init()
	CreateMarkdown("234")
}
