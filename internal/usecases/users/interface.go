package users

import (
	"context"
	"sports-day/internal/entity"
)

type Reader interface {
	Get(ctx context.Context, userID entity.ID) (*entity.User, error)
	GetFromUser(ctx context.Context, userName string) (*entity.User, error)
}
type Writer interface {
	Create(ctx context.Context, u *entity.User) (entity.ID, error)
}

type Repo interface {
	Reader
	Writer
}

type UseCase interface {
	GetUser(ctx context.Context, userID entity.ID) (*entity.User, error)
	GetUserFromUser(ctx context.Context, userName string) (*entity.User, error)
	CreateUser(ctx context.Context, userName, role string) (entity.ID, error)
}
