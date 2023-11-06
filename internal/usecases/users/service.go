package users

import (
	"context"
	"log"
	"sports-day/db"
	"sports-day/internal/entity"
	"sports-day/internal/repository"
	"sync"

	"github.com/google/uuid"
)

var userService UseCase = nil
var serviceOnce sync.Once

func GetUserService() UseCase {
	serviceOnce.Do(func() {
		log.Println("Setting up service")
		db := db.GetDB()
		userRepo := repository.NewUserPg(db)
		userService = newService(userRepo)
	})
	return userService
}

type service struct {
	repo Repo
}

// Newservice creates a new usecase
func newService(r Repo) *service {
	return &service{
		repo: r,
	}
}

func (s *service) GetUser(ctx context.Context, userID entity.ID) (*entity.User, error) {
	return s.repo.Get(ctx, userID)
}
func (s *service) GetUserFromUsername(ctx context.Context, userName string) (*entity.User, error) {
	return s.repo.GetFromUser(ctx, userName)
}
func (s *service) CreateUser(ctx context.Context, userName, role string) (entity.ID, error) {
	u, err := entity.NewUser(userName, role)
	if err != nil {
		return uuid.Nil, err
	}
	return s.repo.Create(ctx, u)
}
