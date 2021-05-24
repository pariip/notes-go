package schema

import (
	"github.com/pariip/notes-go/internal/models"
	"gorm.io/gorm"
)

type (
	Picture struct {
		gorm.Model
		Name   string
		Alt    string
		NoteID uint
	}
)

func (p *Picture) ConvertModel() *models.Picture {
	return &models.Picture{
		ID:     p.ID,
		Name:   p.Name,
		Alt:    p.Alt,
		NoteID: p.NoteID,
	}
}

func ConvertPicture(picture *models.Picture) *Picture {
	return &Picture{
		Model:  gorm.Model{ID: picture.ID},
		Name:   picture.Name,
		Alt:    picture.Alt,
		NoteID: picture.NoteID,
	}
}
