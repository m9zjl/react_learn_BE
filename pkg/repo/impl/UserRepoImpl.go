package impl

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/pkg/domain/entity"
	"server/pkg/repo"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.IUserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

func (u *UserRepoImpl) Add(user *entity.User) (bool, error) {
	ret := u.db.Clauses(clause.Insert{Modifier: "OR IGNORE"}).Create(&user)
	if ret.Error != nil {
		return false, ret.Error
	}
	return ret.RowsAffected > 0, nil
}

func (u *UserRepoImpl) ByEmail(email string) (*entity.User, error) {
	var user = entity.User{}
	ret := u.db.Model(&user).Where("email=?", email).First(&user)
	if ret.Error != nil && errors.Is(ret.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, ret.Error
}

func (u *UserRepoImpl) Remove(id int) error {
	var user = entity.User{
		ID: id,
	}
	return u.db.Model(&user).Delete(&user).Error
}

func (u *UserRepoImpl) GetUserByIds(ids []int) ([]*entity.User, error) {
	var users []*entity.User
	ret := u.db.Where("id in ?", ids).Find(&users)
	if ret.Error != nil && errors.Is(ret.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return users, ret.Error
}
