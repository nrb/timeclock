package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	db, err := NewDB(DB_MODE_MEM)
	require.Nil(t, err)
	us := NewUserStore(db)
	defer db.Close()

	u := &User{
		Name:     "John Cleese",
		Email:    "john@monty.com",
		Password: "pass",
		Salt:     "salt",
		Admin:    false,
	}

	id, err := us.Create(u)

	assert.Equal(t, 1, id)
	assert.Nil(t, err)
}

func TestGetUser(t *testing.T) {
	db, err := NewDB(DB_MODE_MEM)
	require.Nil(t, err)

	us := NewUserStore(db)
	defer db.Close()

	u := &User{
		ID:       1, // We can set ID safely here because the value is ignored by Create. Set it for comparison on the actual test
		Name:     "John Cleese",
		Email:    "john@monty.com",
		Password: "pass",
		Salt:     "salt",
		Admin:    false,
	}
	_, err = us.Create(u)

	u2, err := us.Get(1)

	assert.EqualValues(t, u, u2)
	assert.Nil(t, err)

}
