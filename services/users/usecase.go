package users

import "context"

type Usecase interface {
	Login(ctx context.Context) error
	Register(ctx context.Context) error
}