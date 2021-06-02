package params

type (
	CreateNoteRequest struct {
		UserID      uint   `json:"user_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		PublicNote  bool   `json:"public_note"`
		//Pictures    []Picture
	}
	UpdateNoteRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		PublicNote  bool   `json:"public_note"`
	}
)
