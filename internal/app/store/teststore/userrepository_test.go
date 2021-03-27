package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"http-rest-api/internal/app/store/teststore"

	"testing"
)

const USER_EMAIL = "user@example.org"

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	_, err := s.User().FindByEmail(USER_EMAIL)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = USER_EMAIL
	s.User().Create(u)
	u, err = s.User().FindByEmail(USER_EMAIL)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	u.Email = USER_EMAIL
	s.User().Create(u)
	u, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
