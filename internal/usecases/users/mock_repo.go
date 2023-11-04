package users

import (
	"context"
	"sports-day/internal/entity"
)

// inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.User
}

// newInmem create new repository
func newInmem() *inmem {
	return &inmem{
		m: map[entity.ID]*entity.User{},
	}
}

func (r *inmem) Create(ctx context.Context, u *entity.User) (entity.ID, error) {
	r.m[u.ID] = u
	return u.ID, nil
}

func (r *inmem) Get(ctx context.Context, id entity.ID) (*entity.User, error) {
	u, ok := r.m[id]
	if !ok {
		return nil, entity.ErrUserNotFound
	}
	return u, nil
}
func (r *inmem) GetFromUser(ctx context.Context, userName string) (*entity.User, error) {
	for e, u := range r.m {
		if u.UserName == userName {
			return r.m[e], nil
		}
	}
	return &entity.User{}, entity.ErrUserNotFound
}
