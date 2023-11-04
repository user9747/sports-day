package events

import (
	"context"
	"sports-day/internal/entity"
)

type Reader interface {
	GetAll(ctx context.Context, limit int, page int) ([]entity.Event, error)
	GetAllRegistered(ctx context.Context, limit int, page int, userId entity.ID) ([]entity.Event, error)
}
type Writer interface {
	Register(ctx context.Context, userId entity.ID, eventId int) error
	UnRegister(ctx context.Context, userId entity.ID, eventId int) error
}

type Repo interface {
	Reader
	Writer
}

type UseCase interface {
	GetAllEvents(ctx context.Context, limit int, page int) ([]entity.Event, error)
	GetAllRegisteredEvents(ctx context.Context, limit int, page int, userId entity.ID) ([]entity.Event, error)
	RegisterEvent(ctx context.Context, userId entity.ID, eventId int) error
	UnRegisterEvent(ctx context.Context, userId entity.ID, eventId int) error
}
