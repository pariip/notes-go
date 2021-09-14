package schema

import (
	"github.com/pariip/notes-go/internal/models"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserID      uint `gorm:"not null"`
	Title       string
	Description string
	PublicNote  bool
	Pictures    []Picture
}

func (n *Note) ConvertModel() *models.Note {
	var pictures []models.Picture
	for _, picture := range n.Pictures {
		pictures = append(pictures, models.Picture{
			ID:     picture.ID,
			Name:   picture.Name,
			Alt:    picture.Alt,
			NoteID: picture.NoteID,
		})
	}
	return &models.Note{
		ID:          n.ID,
		UserID:      n.UserID,
		Title:       n.Title,
		Description: n.Description,
		PublicNote:  n.PublicNote,
		Pictures:    pictures,
	}
}
func ConvertNote(note *models.Note) *Note {
	var pictures []Picture
	for _, picture := range note.Pictures {
		pictures = append(pictures, Picture{
			Model:  gorm.Model{ID: picture.ID},
			Name:   picture.Name,
			Alt:    picture.Alt,
			NoteID: picture.NoteID,
		})
	}
	return &Note{
		Model:       gorm.Model{ID: note.ID},
		UserID:      note.UserID,
		Title:       note.Title,
		Description: note.Description,
		PublicNote:  note.PublicNote,
		Pictures:    pictures,
	}
}
