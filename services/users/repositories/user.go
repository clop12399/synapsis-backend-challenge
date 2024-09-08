package repositories

import (
	"synapsis-backend-challenge/services/users"
	"synapsis-backend-challenge/services/users/entities"
)

type UserRepository struct {
}

// StoreUser implements users.Repository.
func (u *UserRepository) StoreUser(user *entities.User) error {
	panic("unimplemented")
}

func NewuserRepository() users.Repository {
	return &UserRepository{}
}
