package contract

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
)

type UserService interface {
	CreateUser(req *params.CreateUserRequest) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(req *params.UpdateUserRequest) (*models.User, error)
	DeleteUser(userID uint) error

	Login(req *params.LoginRequest) (*params.UserTokens, error)
	RefreshToken(refreshToken string, userID uint) (*params.UserTokens, error)
}
