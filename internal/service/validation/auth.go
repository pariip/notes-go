package validation

import (
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"unicode"
)

func (s *service) Username(username string) error {
	if l := len(username); l <= s.cfg.UsernameMinLength || l >= s.cfg.UsernameMaxLength {
		s.logger.Error(&log.Field{
			Section:  "user.validation",
			Function: "validateUsername",
			Params:   map[string]interface{}{"req username": username},
			Message:  s.translator.TranslateEn(messages.InvalidUsernameLength),
		})
		return cerrors.New(cerrors.KindInvalid, messages.InvalidUsernameLength)
	}
	return nil
}

func (s *service) Password(password string) error {
	var number, upper, special bool
	var letters int

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			s.logger.Error(&log.Field{
				Section:  "user.validation",
				Function: "validateCreateUser",
				Params:   map[string]interface{}{"password": password},
				Message:  s.translator.TranslateEn(messages.InvalidPassword),
			})

			return cerrors.New(cerrors.KindInvalid, messages.InvalidPassword)
		}
	}
	if letters >= s.cfg.PasswordMinLetters && number && upper && special {
		return nil
	}
	s.logger.Error(&log.Field{
		Section:  "user.validation",
		Function: "validateCreateUser",
		Params:   map[string]interface{}{"password": password},
		Message:  s.translator.TranslateEn(messages.InvalidPassword),
	})

	return cerrors.New(cerrors.KindInvalid, messages.InvalidPassword)
}
