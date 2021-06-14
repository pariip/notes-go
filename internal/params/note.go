package params

import "github.com/pariip/notes-go/internal/models"

type (
	CreateNoteRequest struct {
		UserID      uint             `json:"user_id"`
		Title       string           `json:"title"`
		Description string           `json:"description"`
		PublicNote  bool             `json:"public_note"`
		Pictures    []models.Picture `json:"pictures"`
	}
	UpdateNoteRequest struct {
		ID          uint
		Title       string           `json:"title"`
		Description string           `json:"description"`
		PublicNote  bool             `json:"public_note"`
		Pictures    []models.Picture `json:"pictures"`
	}
)
