package note

import (
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (s *service) validation(note *params.CreateNoteRequest) error {
	if note.Title == "" {
		s.logger.Error(&log.Field{
			Section:  "note.validation",
			Function: "validation",
			Params:   map[string]interface{}{"req note": note},
			Message:  s.translator.TranslateEn(messages.InvalidEmptyTitle),
		})
		return cerrors.New(cerrors.KindInvalid, messages.InvalidEmptyTitle)
	}
	if note.Description == "" {
		s.logger.Error(&log.Field{
			Section:  "note.validation",
			Function: "validation",
			Params:   map[string]interface{}{"req note": note},
			Message:  s.translator.TranslateEn(messages.InvalidEmptyDescription),
		})
		return cerrors.New(cerrors.KindInvalid, messages.InvalidEmptyDescription)
	}

	return nil
}
