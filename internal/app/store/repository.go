package store

import "drawserver/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByName(string) (*model.User, error)
}
