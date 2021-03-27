package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"http-rest-api/internal/app/store/sqlstore"
	"testing"
)

const USER_EMAIL = "user@example.org"

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
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
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)
	u.Email = USER_EMAIL
	s.User().Create(u)
	u, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
