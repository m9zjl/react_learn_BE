package repo

import (
	"server/pkg/domain/dto"
	"server/pkg/domain/entity"
)

type IArticleRepo interface {
	Save(article *entity.Article) (bool, error)
	Remove(id int) error
	Query(*dto.ArticleQueryParameter) ([]*entity.Article, error)
}
