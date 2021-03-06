package postgres

import (
	"errors"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/random"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"math/rand"
	"testing"
)

func TestCreateUser(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()

	tests := []struct {
		name string
		user *models.User
		want error
	}{
		{
			name: "create user",
			user: user,
			want: nil,
		},
		{
			name: "username is not unique",
			user: user,
			want: cerrors.New(cerrors.KindUnexpected, messages.DBError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repositoryTest.CreateUser(tt.user)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	var err error

	t.Run("create new test user", func(t *testing.T) {
		if user, err = repositoryTest.CreateUser(user); err != nil {
			t.Fail()
		}
	})

	tests := []struct {
		name   string
		userID uint
		want   error
	}{
		{
			name:   "get user by id",
			userID: user.ID,
			want:   err,
		},
		{
			name:   "user not found",
			userID: uint(rand.Uint64()),
			want:   cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repositoryTest.GetUserByID(tt.userID)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	var err error

	t.Run("create new test user", func(t *testing.T) {
		if user, err = repositoryTest.CreateUser(user); err != nil {
			t.Fail()
		}
	})

	tests := []struct {
		name     string
		username string
		want     error
	}{
		{
			name:     "get user by id",
			username: user.Username,
			want:     err,
		},
		{
			name:     "user not found",
			username: random.String(10),
			want:     cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repositoryTest.GetUserByUsername(tt.username)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	user1 := newUserTest()
	var err error

	t.Run("create new test user", func(t *testing.T) {
		if user, err = repositoryTest.CreateUser(user); err != nil {
			t.Fail()
		}
	})
	tests := []struct {
		name string
		user *models.User
		want error
	}{
		{
			name: "update user",
			user: user,
			want: nil,
		},
		{
			name: "user not found",
			user: user1,
			want: cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repositoryTest.UpdateUser(tt.user)
			if !errors.Is(err, tt.want) {
				t.Error()
			}

		})
	}
}

func TestDeleteUser(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	user1 := newUserTest()
	var err error

	t.Run("create new record", func(t *testing.T) {
		if user, err = repositoryTest.CreateUser(user); err != nil {
			t.Fail()
		}
	})
	user1.ID = uint(rand.Uint32())

	tests := []struct {
		name string
		user *models.User
		want error
	}{
		// TODO: Add test cases.
		{
			name: "delete user",
			user: user,
			want: nil,
		},
		{
			name: "user not found",
			user: user1,
			want: cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repositoryTest.DeleteUser(tt.user)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}
