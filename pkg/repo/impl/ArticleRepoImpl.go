package impl

import (
	"errors"
	"github.com/duke-git/lancet/v2/strutil"
	"gorm.io/gorm"
	"server/pkg/domain/dto"
	"server/pkg/domain/entity"
	"server/pkg/repo"
)

type ArticleRepoImpl struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) repo.IArticleRepo {
	return &ArticleRepoImpl{
		db: db,
	}
}

func (a *ArticleRepoImpl) Save(article *entity.Article) (bool, error) {
	var ret *gorm.DB
	if article.ID == 0 {
		ret = a.db.Create(article)
	} else {
		ret = a.db.Updates(article)
	}

	if ret.Error != nil {
		return false, ret.Error
	}
	return ret.RowsAffected > 0, nil

}

func (a *ArticleRepoImpl) Remove(id int) error {
	var article = entity.Article{
		ID: id,
	}
	return a.db.Model(&article).Delete(&article).Error
}

func (a *ArticleRepoImpl) Query(parameter *dto.ArticleQueryParameter) ([]*entity.Article, error) {

	var articles []*entity.Article
	db := a.db
	if parameter.Id != 0 {
		db = db.Where("id = ?", parameter.Id)
	}
	if strutil.IsNotBlank(parameter.Category) {
		db = db.Where("category = ?", parameter.Category)
	}
	ret := db.Find(&articles)
	if ret.Error != nil && errors.Is(ret.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return articles, ret.Error

}
