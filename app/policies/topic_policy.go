package policies

import (
	"github.com/flyrory/go-web-blog/app/models/article"
	"github.com/flyrory/go-web-blog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}
