package contract

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
)

type UserService interface {
	CreateUser(req *params.CreateUserRequest) (*models.User, error)
}
