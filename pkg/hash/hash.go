package hash

import (
	"errors"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"golang.org/x/crypto/bcrypt"
)

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", cerrors.New(cerrors.KindUnexpected, messages.GeneralError)
	}
	return string(bytes), nil
}

func CheckPassword(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return cerrors.New(cerrors.KindInvalid, messages.UsernameOrPasswordIsIncorrect)
		}
		return cerrors.New(cerrors.KindUnexpected, messages.GeneralError)
	}
	return nil
}
