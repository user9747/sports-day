package repository

import (
	"context"
	"log"
	"sports-day/internal/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type EventsPg struct {
	db *sqlx.DB
}

// GetAll implements events.Repo.
func (pg *EventsPg) GetAll(ctx context.Context, limit int, page int) ([]entity.Event, error) {
	var err error
	var uList []entity.Event
	params := make(map[string]interface{})
	var pagination string
	if limit > 0 && page > 0 {
		offset := (page - 1) * limit
		pagination = ` LIMIT :limit OFFSET :offset `
		params["offset"] = offset
		params["limit"] = limit
	}
	query := `SELECT * FROM events order by created_at` + pagination
	namedQuery, err := pg.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	err = namedQuery.SelectContext(ctx, &uList, params)
	return uList, err
}

// GetAllRegistered implements events.Repo.
func (pg *EventsPg) GetAllRegistered(ctx context.Context, limit int, page int, userId uuid.UUID) ([]entity.Event, error) {
	var err error
	var uList []entity.Event
	params := make(map[string]interface{})
	params["userId"] = userId
	var pagination string
	if limit > 0 && page > 0 {
		offset := (page - 1) * limit
		pagination = ` LIMIT :limit OFFSET :offset `
		params["offset"] = offset
		params["limit"] = limit
	}
	query := `SELECT e.* FROM events e 
				JOIN registrations r on e.id = r.event_id where r.user_id=:userId ` + pagination
	namedQuery, err := pg.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	err = namedQuery.SelectContext(ctx, &uList, params)
	return uList, err
}

// Register implements events.Repo.
func (pg *EventsPg) Register(ctx context.Context, userId entity.ID, eventId int) error {
	query := "INSERT INTO registrations(user_id, event_id) VALUES($1, $2);"
	_, err := pg.db.ExecContext(ctx, query, userId, eventId)
	return err
}

// UnRegister implements events.Repo.
func (pg *EventsPg) UnRegister(ctx context.Context, userId entity.ID, eventId int) error {
	query := "DELETE FROM registrations where user_id = $1 and event_id = $2"
	log.Println(query, userId, eventId)
	_, err := pg.db.ExecContext(ctx, query, userId, eventId)
	return err
}

func NewEventsPg(db *sqlx.DB) *EventsPg {
	return &EventsPg{
		db: db,
	}
}
