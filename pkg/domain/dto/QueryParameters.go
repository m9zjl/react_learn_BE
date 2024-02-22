package dto

type ArticleQueryParameter struct {
	Category string `json:"category"`
	Id       int    `json:"id"`
	PageSize int    `json:"pageSize"`
	PageNum  int    `json:"pageNum"`
}
