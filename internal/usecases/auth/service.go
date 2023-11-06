package auth

import (
	"context"
	"database/sql"
	"log"
	"sports-day/internal/entity"
	"sports-day/internal/repository"
	"sports-day/internal/usecases/users"
	"sports-day/internal/utils"
	"sync"
)

var authService UseCase = nil
var serviceOnce sync.Once

func GetAuthService() UseCase {
	serviceOnce.Do(func() {
		log.Println("Setting up Auth service ...")

		u := users.GetUserService()
		r := repository.NewAuthRepo()
		authService = newService(r, u)
	})
	return authService
}

type service struct {
	repo  Repo
	users users.UseCase
}

// NewService creates a new usecase
func newService(r Repo, u users.UseCase) *service {
	return &service{
		repo:  r,
		users: u,
	}
}

// GetUserFromToken implements UseCase.
func (s *service) GetUserFromToken(ctx context.Context, token string) (*entity.LoggedInUser, error) {
	return s.repo.GetUser(ctx, token)
}

// Login implements UseCase.
func (s *service) Login(ctx context.Context, userName string) (string, error) {
	u, err := s.users.GetUserFromUsername(ctx, userName)
	if err == sql.ErrNoRows {
		return "", entity.ErrUserNotFound
	}
	if err != nil {
		return "", err
	}
	token := utils.GenerateRandomString(24)
	err = s.repo.Login(ctx, u, token)
	return token, err
}

// Logout implements UseCase.
func (s *service) Logout(ctx context.Context, token string) error {
	return s.repo.Logout(ctx, token)
}
