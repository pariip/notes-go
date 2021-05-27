package schema

import (
	"github.com/pariip/notes-go/internal/models"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublicNote  bool   `json:"public_note"`
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
