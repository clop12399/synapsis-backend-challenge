package users

import "synapsis-backend-challenge/services/users/entities"

type Repository interface {
	StoreUser(user *entities.User) error
}