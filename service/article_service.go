package service

import (
	"errors"
	"server/pkg/domain/dto"
	"server/pkg/domain/entity"
	"server/pkg/repo"
	"time"
)

type ArticleService struct {
	articleRepo repo.IArticleRepo
	userRepo    repo.IUserRepo
}

func NewArticleService(
	articleRepo repo.IArticleRepo,
	userRepo repo.IUserRepo,

) *ArticleService {
	return &ArticleService{
		articleRepo: articleRepo,
		userRepo:    userRepo,
	}
}

func (a *ArticleService) QueryArticles(queryParameters *dto.ArticleQueryParameter) ([]*dto.ArticleDTO, error) {
	articles, err := a.articleRepo.Query(queryParameters)
	if err != nil {
		return nil, err
	}

	userIds := make([]int, 0)

	for _, article := range articles {
		userIds = append(userIds, article.UserId)
	}

	users, err := a.userRepo.GetUserByIds(userIds)
	if err != nil {
		return nil, err
	}

	userMap := make(map[int]*entity.User)

	for _, user := range users {
		userMap[user.ID] = user
	}

	articleDTOs := make([]*dto.ArticleDTO, 0)

	for _, article := range articles {
		user, _ := userMap[article.UserId]

		tmp := &dto.ArticleDTO{
			Article: article,
			User:    user,
		}

		articleDTOs = append(articleDTOs, tmp)
	}

	return articleDTOs, nil
}

func (a *ArticleService) Add(articleDTO *dto.AddArticleDTO) (bool, error) {
	if articleDTO == nil || articleDTO.User == nil {
		return false, errors.New("article or user is nil")
	}
	article := &entity.Article{
		ID:          articleDTO.ID,
		Title:       articleDTO.Title,
		Desc:        articleDTO.Desc,
		Img:         articleDTO.Img,
		Uid:         articleDTO.User.ID,
		UserId:      articleDTO.User.ID,
		Category:    articleDTO.Category,
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
	}
	return a.articleRepo.Save(article)
}
