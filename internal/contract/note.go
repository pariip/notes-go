package contract

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
)

//go:generate mockgen -source ./note.go  -package note_mock  -destination ../mock/note_mock/mock.go

type NoteService interface {
	CreateNote(req *params.CreateNoteRequest) (*models.Note, error)
	GetAllNotes(userID uint) ([]*models.Note, error)
	GetNoteByID(noteID uint) (*models.Note, error)
	UpdateNote(req *params.UpdateNoteRequest) (*models.Note, error)
	DeleteNote(noteID uint) error
}
