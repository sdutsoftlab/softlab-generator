package gen

type Website struct {
	Title      string `yaml:"title"`
	WebsiteUrl string `yaml:"url"`

	SubTitle    string `yaml:"subtitle"`
	Description string `yaml:"description"`
	Keywords    string `yaml:"keywords"`

	Markdown string `yaml:"markdown"` // md存储路径
	Dist     string `yaml:"output"`   // 输出编译输出文件

	Theme string `yaml:"theme"`

	Author string `yaml:"author"`
	Avatar string `yaml:"avatar"`
	Github string `yaml:"github"`
	Mail   string `yaml:"mail"`

	Cates []string `yaml:"cate"`
	Paths []string `yaml:"paths"`
	Exts  []string `yaml:"exts"`

	HomeArtNum   int    `yaml:"pageNum"`
	HomeTitle    string `yaml:"home_title,omitempty"`
	ArchiveTitle string `yaml:"archive_title,omitempty"`
	TagTitle     string `yaml:"tag_title,omitempty"`
	CateTitle    string `yaml:"cate_title,omitempty"`
	AboutTitle   string `yaml:"about_title,omitempty"`
	ArticleTitle string `yaml:"article_title,omitempty"`
}

// Article 定义了一篇文章所需要的要素
type Article struct {
	Id          int    `json:"article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	Summary  string   `json:"summary"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
	Category []string `json:"cate"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Url       string `json:"url"`
}
