package contract

import "github.com/pariip/notes-go/internal/params"

//go:generate mockgen -source ./validation.go -package validation_mock -destination ../mock/validation_mock/mock.go

type (
	ValidationService interface {
		AuthValidation
		NoteValidation
	}
	AuthValidation interface {
		Username(username string) error
		Password(password string) error
	}
	NoteValidation interface {
		Note(note *params.CreateNoteRequest) error
	}
)
