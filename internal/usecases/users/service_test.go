package users

import (
	"context"
	"sports-day/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateInvalidDashboardUser(t *testing.T) {
	r := newInmem()
	s := newService(r)
	_, err := s.CreateUser(context.Background(), "", "admin")
	assert.ErrorIs(t, err, entity.ErrInvalidUserName)
	_, err = s.CreateUser(context.Background(), "testuser", "invalidrole")
	assert.ErrorIs(t, err, entity.ErrInvalidRole)
}

func Test_CreateDashboardUser(t *testing.T) {
	r := newInmem()
	s := newService(r)
	_, err := s.CreateUser(context.Background(), "testuser", "admin")
	if err != nil {
		t.Error("dashboard user creation failed", err)
	}
}

func Test_GetDashboardUser(t *testing.T) {
	r := newInmem()
	s := newService(r)
	userID, err := s.CreateUser(context.Background(), "testuser", "admin")
	assert.Nil(t, err, "User creation failed")

	u, err := s.GetUser(context.Background(), userID)
	assert.Nil(t, err, "Unable to get user")
	assert.Equal(t, "testuser", u.UserName, "username not matching")
	assert.Equal(t, u.ID, userID, "userID not matching")
}

func Test_GetDashboardUserWithUserName(t *testing.T) {
	r := newInmem()
	s := newService(r)
	userName := "testuser"
	userID, err := s.CreateUser(context.Background(), userName, "admin")
	assert.Nil(t, err, "User creation failed")

	u, err := s.GetUserFromUsername(context.Background(), userName)
	assert.Nil(t, err, "Unable to get user")
	assert.Equal(t, userName, u.UserName, "username not matching")
	assert.Equal(t, u.ID, userID, "userID not matching")
}
