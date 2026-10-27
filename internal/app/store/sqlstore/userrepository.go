package sqlstore

import (
	"drawserver/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (name, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Name,
		u.EncryptedPassword,
	).Scan(&u.Id)
}

func (r *UserRepository) FindByName(name string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, name, encrypted_password FROM users WHERE name = $1",
		name,
	).Scan(
		&u.Id,
		&u.Name,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}
