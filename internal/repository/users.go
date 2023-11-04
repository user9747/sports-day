package repository

import (
	"context"
	"sports-day/internal/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserPg struct {
	db *sqlx.DB
}

func NewUserPg(db *sqlx.DB) *UserPg {
	return &UserPg{
		db: db,
	}
}

// Create implements users.Repo.
func (pg *UserPg) Create(ctx context.Context, u *entity.User) (uuid.UUID, error) {
	query := "INSERT INTO users(id, username, role) VALUES($1, $2, $3);"
	var err error
	_, err = pg.db.ExecContext(ctx, query, u.ID, u.UserName, u.Role)
	return u.ID, err
}

// Get implements users.Repo.
func (pg *UserPg) Get(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	query := `SELECT * from users where id = $1;`
	var err error
	var u entity.User
	err = pg.db.GetContext(ctx, &u, query, userID)
	return &u, err
}

// GetFromUser implements users.Repo.
func (pg *UserPg) GetFromUser(ctx context.Context, userName string) (*entity.User, error) {
	query := `SELECT * from users where username = $1;`
	var err error
	var u entity.User
	err = pg.db.GetContext(ctx, &u, query, userName)
	return &u, err
}
