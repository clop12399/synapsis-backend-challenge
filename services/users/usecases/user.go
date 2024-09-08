package usecases

import (
	"context"
	"synapsis-backend-challenge/services/users"
)

type UserUsecase struct {
	repo users.Repository
}

func NewUserUsecase(repo users.Repository) users.Usecase {
	return &UserUsecase{
		repo: repo,
	}
}

// Login implements users.Usecase.
func (u *UserUsecase) Login(ctx context.Context) error {
	panic("unimplemented")
}

// Register implements users.Usecase.
func (u *UserUsecase) Register(ctx context.Context) error {
	panic("unimplemented")
}