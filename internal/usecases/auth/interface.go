package auth

import (
	"context"
	"sports-day/internal/entity"
)

type Reader interface {
	GetUser(ctx context.Context, token string) (*entity.LoggedInUser, error)
}
type Writer interface {
	Login(ctx context.Context, u *entity.User, token string) error
	Logout(ctx context.Context, token string) error
}

type Repo interface {
	Reader
	Writer
}

type UseCase interface {
	Login(ctx context.Context, userName string) (string, error)
	Logout(ctx context.Context, token string) error
	GetUserFromToken(ctx context.Context, token string) (*entity.LoggedInUser, error)
}
