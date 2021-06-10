package contract

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
)

//go:generate mockgen -source ./auth.go  -package auth_mock  -destination ../mock/auth_mock/mock.go

type AuthService interface {
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
	RefreshToken(refreshToken string, userID uint) (*params.UserTokens, error)

	Login(req *params.LoginRequest) (*params.UserTokens, error)
	Signup(req *params.SignupRequest) (*params.UserTokens, error)
}
