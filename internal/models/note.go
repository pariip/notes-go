package models

type Note struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublicNote  bool   `json:"public_note"`
	Pictures    []Picture
}
