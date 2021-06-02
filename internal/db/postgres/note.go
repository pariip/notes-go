package postgres

import (
	"github.com/pariip/notes-go/internal/db/postgres/schema"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (r *repository) CreateNote(note *models.Note) (*models.Note, error) {
	n := schema.ConvertNote(note)
	if err := r.db.Create(n).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"user": n},
			Message:  err.Error(),
		})
		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	return n.ConvertModel(), nil
}

func (r *repository) GetAllNotes(userID uint) ([]*models.Note, error) {
	notes := make([]schema.Note, 0)

	if err := r.db.Model(&schema.Note{}).Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.note",
			Function: "GetAllNotes",
			Params:   map[string]interface{}{"user_id": userID},
			Message:  err.Error(),
		})
		if isErrorNotFound(err) {
			return nil, cerrors.New(cerrors.KindNotFound, messages.NoteNotFound)
		}

		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	repNote := make([]*models.Note, 0)
	for _, note := range notes {
		repNote = append(repNote, note.ConvertModel())
	}

	return repNote, nil
}

func (r *repository) GetNoteByID(noteID uint) (*models.Note, error) {
	note := new(schema.Note)

	if err := r.db.Model(&schema.Note{}).Where("id = ?", noteID).First(note).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.note",
			Function: "GetNoteById",
			Params:   map[string]interface{}{"noteID": noteID},
			Message:  err.Error(),
		})
		if isErrorNotFound(err) {
			return nil, cerrors.New(cerrors.KindNotFound, messages.NoteNotFound)

		}
		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}

	return note.ConvertModel(), nil
}

func (r *repository) UpdateNote(note *models.Note) (*models.Note, error) {
	n := schema.ConvertNote(note)

	if err := r.db.Model(&schema.Note{}).First(&schema.Note{}, n.ID).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.note",
			Function: "UpdateNote",
			Params:   map[string]interface{}{"note": n},
			Message:  err.Error(),
		})

		if isErrorNotFound(err) {
			return nil, cerrors.New(cerrors.KindNotFound, messages.NoteNotFound)
		}

		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}

	if err := r.db.Model(&schema.Note{}).Where("id =?", n.ID).Save(n).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.note",
			Function: "UpdateUser",
			Params:   map[string]interface{}{"note": n},
			Message:  err.Error(),
		})
		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	return n.ConvertModel(), nil
}
