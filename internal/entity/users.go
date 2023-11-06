package entity

import "errors"

type User struct {
	ID        ID     `db:"id"`
	UserName  string `db:"username"`
	Role      string `db:"role"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	CreatedBy ID     `db:"created_by"`
}

type LoggedInUser struct {
	ID       ID     `db:"id"`
	UserName string `db:"username"`
	Role     string `db:"role"`
}

const (
	UserRoleParticipant = "participant"
	UserRoleAdmin       = "admin"
)

var (
	ErrInvalidRole     = errors.New("invalid role")
	ErrInvalidUserName = errors.New("username length should be between 4 and 20")
	ErrUserNotFound    = errors.New("user not found")
)

func NewUser(userName, role string) (*User, error) {
	if len(userName) < 4 || len(userName) > 20 {
		return nil, ErrInvalidUserName
	}
	if err := validateRole(role); err != nil {
		return nil, err
	}
	u := &User{
		ID:       NewID(),
		UserName: userName,
		Role:     role,
	}

	return u, nil
}

func validateRole(role string) error {
	if role != UserRoleParticipant && role != UserRoleAdmin {
		return ErrInvalidRole
	}
	return nil
}
