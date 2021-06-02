package note

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
)

func (s service) CreateNote(req *params.CreateNoteRequest) (*models.Note, error) {
	if err := s.validation(req); err != nil {
		return nil, err
	}

	note := &models.Note{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		PublicNote:  req.PublicNote,
		Pictures:    nil,
	}

	note, err := s.noteRepo.CreateNote(note)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (s service) GetAllNotes(userID uint) ([]*models.Note, error) {
	var notes []*models.Note

	notes, err := s.noteRepo.GetAllNotes(userID)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s service) GetNoteByID(noteId uint) (*models.Note, error) {
	note, err := s.noteRepo.GetNoteByID(noteId)

	if err != nil {
		return nil, err
	}
	return note, nil
}
