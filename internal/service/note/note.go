package note

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (s service) CreateNote(req *params.CreateNoteRequest) (*models.Note, error) {
	if err := s.validate.Note(req); err != nil {
		return nil, err
	}

	note := &models.Note{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		PublicNote:  req.PublicNote,
	}
	if req.Pictures != nil {
		var pic []models.Picture
		for _, picture := range req.Pictures {
			photo, err := s.noteRepo.IsImageExist(picture.Alt)
			if err != nil {
				return nil, err
			}
			if _, err = imageIsDuplicate(photo, picture.ID); err != nil {
				s.logger.Error(&log.Field{
					Section:  "service.note",
					Function: "CreateNote",
					Params:   map[string]interface{}{"pictureName": photo.Name},
					Message:  s.translator.Translate(messages.PictureDuplicate),
				})
				return nil, err
			}
			picture = *photo
			pic = append(pic, picture)
		}
		note.Pictures = pic
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

func (s service) GetAllMyNotes(userID uint) ([]*models.Note, error) {
	var notes []*models.Note

	notes, err := s.noteRepo.GetAllMyNotes(userID)
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

func (s service) UpdateNote(req *params.UpdateNoteRequest) (*models.Note, error) {
	note, err := s.noteRepo.GetNoteByID(req.ID)
	if err != nil {
		return nil, err
	}
	note.Title = req.Title
	note.Description = req.Description
	note.PublicNote = req.PublicNote

	note, err = s.noteRepo.UpdateNote(note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s service) DeleteNote(noteID uint) error {
	note, err := s.noteRepo.GetNoteByID(noteID)
	if err != nil {
		return err
	}
	if err := s.noteRepo.DeleteNote(note); err != nil {
		return err
	}
	return nil
}
