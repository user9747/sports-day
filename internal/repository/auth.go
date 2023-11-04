package repository

import (
	"context"
	"encoding/json"
	"sports-day/cache"
	"sports-day/internal/entity"
	"time"
)

type AuthRepo struct {
}

// GetUser implements auth.Repo.
func (*AuthRepo) GetUser(ctx context.Context, token string) (*entity.User, error) {
	userObjString, err := cache.Get(ctx, token)
	if err != nil {
		return nil, err
	}
	var u entity.User
	err = json.Unmarshal([]byte(userObjString), &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Login implements auth.Repo.
func (*AuthRepo) Login(ctx context.Context, u *entity.User, token string) error {
	return cache.SetStruct(token, *u, time.Hour*24)
}

// Logout implements auth.Repo.
func (*AuthRepo) Logout(ctx context.Context, token string) error {
	return cache.Delete(ctx, token)
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}
