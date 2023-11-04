package events

import (
	"context"
	"log"
	"sports-day/db"
	"sports-day/internal/entity"
	"sports-day/internal/repository"
	"sync"
)

var serviceOnce sync.Once
var eventsService UseCase = nil

func GetService() UseCase {
	serviceOnce.Do(func() {
		log.Println("Setting up Dashboard user service")
		db := db.GetDB()
		eventsRepo := repository.NewEventsPg(db)
		eventsService = newService(eventsRepo)
	})
	return eventsService
}

type service struct {
	repo Repo
}

// newService creates a new usecase
func newService(r Repo) *service {
	return &service{
		repo: r,
	}
}

func (s *service) GetAllEvents(ctx context.Context, limit int, page int) ([]entity.Event, error) {
	return s.repo.GetAll(ctx, limit, page)
}
func (s *service) GetAllRegisteredEvents(ctx context.Context, limit int, page int, userID entity.ID) ([]entity.Event, error) {
	return s.repo.GetAllRegistered(ctx, limit, page, userID)
}
func (s *service) RegisterEvent(ctx context.Context, userId entity.ID, eventId int) error {
	return s.repo.Register(ctx, userId, eventId)
}
func (s *service) UnRegisterEvent(ctx context.Context, userId entity.ID, eventId int) error {
	return s.repo.UnRegister(ctx, userId, eventId)
}
