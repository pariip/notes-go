package validation

import (
	"errors"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/random"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	tests := []struct {
		name     string
		username string
		want     error
	}{
		{
			name:     "valid username",
			username: "pari",
			want:     nil,
		},
		{
			name:     "invalid username",
			username: random.String(2),
			want:     cerrors.New(cerrors.KindInvalid, messages.InvalidUsernameLength),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := serviceTest.Username(tt.username)
			if !errors.Is(tt.want, err) {
				t.Fail()
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	test := []struct {
		name     string
		password string
		want     error
	}{
		{
			name:     "correct password",
			password: "Bookstore&*2021",
			want:     nil,
		},
		{
			name:     "invalid password for upper letter",
			password: "parsa%&52",
			want:     cerrors.New(cerrors.KindInvalid, messages.InvalidPassword),
		},
		{
			name:     "invalid password for symbol",
			password: "CdaVddd626",
			want:     cerrors.New(cerrors.KindInvalid, messages.InvalidPassword),
		},
		{
			name:     "invalid password for letter",
			password: "15114339&%@621848",
			want:     cerrors.New(cerrors.KindInvalid, messages.InvalidPassword),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := serviceTest.Password(tt.password)
			if !errors.Is(err, tt.want) {
				t.Fail()
			}
		})
	}
}
