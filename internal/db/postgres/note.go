package postgres

import (
	"fmt"
	"github.com/pariip/notes-go/internal/db/postgres/schema"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/models/types"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

//CreateNote create new note for current user
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

//GetAllNotes get all notes in database with this rule:
/* if user Role is admin or top he/she can show all notes in database(private and public)
else user role is basic can see only her/his notes and only public notes */
func (r *repository) GetAllNotes(userID uint) ([]*models.Note, error) {
	notes := make([]schema.Note, 0)

	user, err := r.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	fmt.Println(userID, user)
	if user.Role == types.Admin || user.Role == types.Top {
		if err := r.db.Model(&schema.Note{}).Where("public_note =? or public_note=? ", false, true).Find(&notes).Error; err != nil {
			r.logger.Error(&log.Field{
				Section:  "repository.note",
				Function: "GetAllNotes",
				Params:   map[string]interface{}{"user_id": userID},
				Message:  err.Error(),
			})
			return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
		}

	} else {
		if err := r.db.Model(&schema.Note{}).Where("user_id = ? or public_note=?", userID, true).Find(&notes).Error; err != nil {
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
	}

	repNote := make([]*models.Note, 0)
	for _, note := range notes {
		repNote = append(repNote, note.ConvertModel())
	}

	return repNote, nil
}

//GetAllMyNotes get all notes for current user
func (r *repository) GetAllMyNotes(userID uint) ([]*models.Note, error) {
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

//GetNoteByID get an note with noteID
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

//UpdateNote user can update title or description or public_note
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

//DeleteNote can delete note with especial ID
func (r *repository) DeleteNote(note *models.Note) error {
	n := schema.ConvertNote(note)

	res := r.db.Model(&schema.Note{}).Where("id = ?", n.ID).Delete(n)
	if err := res.Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.note",
			Function: "DeleteUser",
			Params:   map[string]interface{}{"note": n},
			Message:  err.Error(),
		})
		return cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	if res.RowsAffected != 1 {
		r.logger.Error(&log.Field{
			Section:  "repository.note",
			Function: "DeleteUser",
			Params:   map[string]interface{}{"note": n},
			Message:  r.translator.Translate(messages.NoteNotFound),
		})
		return cerrors.New(cerrors.KindNotFound, messages.NoteNotFound)
	}
	return nil
}
