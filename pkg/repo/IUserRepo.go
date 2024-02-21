package repo

import "server/pkg/domain/entity"

type IUserRepo interface {
	Add(user *entity.User) (bool, error)
	Remove(id int) error
	ByEmail(email string) (*entity.User, error)
}
