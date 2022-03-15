package category

import "github.com/flyrory/go-web-blog/app/models"

// Category 文章分类
type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}
