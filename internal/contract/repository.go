package contract

import "github.com/pariip/notes-go/internal/models"

//go:generate mockgen -source ./repository.go -package repository_mock -destination ../mock/repository_mock/mock.go

type (
	MainRepository interface {
		UserRepository
		NoteRepository
		ImageRepository
	}
	UserRepository interface {
		CreateUser(user *models.User) (*models.User, error)
		GetUserByID(id uint) (*models.User, error)
		GetUserByUsername(username string) (*models.User, error)
		UpdateUser(user *models.User) (*models.User, error)
		DeleteUser(user *models.User) error
		IsUsernameExist(username string) (bool, error)
	}

	NoteRepository interface {
		CreateNote(note *models.Note) (*models.Note, error)
		GetAllNotes(userID uint) ([]*models.Note, error)
		GetAllMyNotes(userID uint) ([]*models.Note, error)
		GetNoteByID(id uint) (*models.Note, error)
		UpdateNote(note *models.Note) (*models.Note, error)
		DeleteNote(note *models.Note) error
	}

	ImageRepository interface {
		UploadImage(pic *models.Picture) (*models.Picture, error)
		IsImageExist(picAlt string) (*models.Picture, error)
	}
)
