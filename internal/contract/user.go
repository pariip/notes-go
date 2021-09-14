package contract

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
)

//go:generate mockgen -source ./user.go -package user_mock -destination ../mock/user_mock/mock.go

type UserService interface {
	CreateUser(req *params.CreateUserRequest) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(req *params.UpdateUserRequest) (*models.User, error)
	DeleteUser(userID uint) error
}
