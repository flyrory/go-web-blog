package article

import (
	"github.com/flyrory/go-web-blog/app/models"
	"github.com/flyrory/go-web-blog/pkg/route"
	"github.com/flyrory/go-web-blog/pkg/types"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", types.Uint64ToString(a.ID))
}
