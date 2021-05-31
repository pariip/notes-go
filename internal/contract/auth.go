package contract

import "github.com/pariip/notes-go/internal/models"

type AuthService interface {
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
	RefreshTokenIsValid(token string, userID uint) (bool, error)
}
