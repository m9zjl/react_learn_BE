package dto

import "server/pkg/domain/entity"

type ArticleDTO struct {
	*entity.Article
	User *entity.User `json:"user"`
}

type AddArticleDTO struct {
	ID       int          `json:"id"`
	Title    string       `json:"title"`
	Desc     string       `json:"desc"`
	User     *entity.User `json:"user"`
	Img      string       `json:"img"`
	Category string       `json:"category"`
}
