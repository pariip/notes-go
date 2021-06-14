package note

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func imageIsDuplicate(photo *models.Picture, noteID uint) (bool, error) {
	if photo.NoteID != nil && photo.NoteID != &noteID {
		return true, cerrors.New(cerrors.KindUnexpected, messages.PictureDuplicate)
	}
	return false, nil
}
