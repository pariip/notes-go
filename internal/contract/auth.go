package contract

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
)

type AuthService interface {
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
	RefreshToken(refreshToken string, userID uint) (*params.UserTokens, error)

	Login(req *params.LoginRequest) (*params.UserTokens, error)
}
